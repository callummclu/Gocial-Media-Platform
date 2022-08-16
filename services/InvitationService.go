package services

import (
	"errors"
	"fmt"

	"github.com/callummclu/Gocial-Media-Platform/auth"
	"github.com/callummclu/Gocial-Media-Platform/configs"
	"github.com/callummclu/Gocial-Media-Platform/middleware"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

func SendUserInvitation(username string, sentUsername string, token string) error {

	err := auth.CheckJWT(token, &username)

	if err != nil {
		return errors.New("Invalid Token")
	}

	db, err := configs.GetDB()
	if err != nil {
		fmt.Print("DB ERRRO")
		err = errors.New("DB connection error")
		return err
	}
	defer db.Close()

	// Check request isnt already sent/received

	var username_friends []string

	username_friends_stmt := "SELECT friends from users WHERE username=$1"

	if err := db.QueryRow(username_friends_stmt, username).Scan(pq.Array(&username_friends)); err != nil {
		return errors.New("Failed to get username requests")
	}

	if middleware.Contains(username_friends, sentUsername) {
		return errors.New("you are already friends with this user")
	}

	var username_received_requests []string

	username_received_stmt := "SELECT received_invitations from users WHERE username=$1"

	if err := db.QueryRow(username_received_stmt, username).Scan(pq.Array(&username_received_requests)); err != nil {
		return errors.New("Failed to get username requests")
	}

	if middleware.Contains(username_received_requests, sentUsername) {
		fmt.Println("request already received")

		return errors.New("No need to send already received")
	}

	var username_sent_requests []string
	var sentUsername_received_requests []string

	username_stmt := "SELECT sent_invitations from users WHERE username=$1"
	sentUsername_stmt := "SELECT received_invitations from users WHERE username=$1"

	if err := db.QueryRow(username_stmt, username).Scan(pq.Array(&username_sent_requests)); err != nil {
		fmt.Print("username requests")
		return errors.New("Failed to get username requests")

	}

	if middleware.Contains(username_sent_requests, sentUsername) {
		fmt.Println("request already sent")
		return errors.New("request already sent")
	}

	if err := db.QueryRow(sentUsername_stmt, sentUsername).Scan(pq.Array(&sentUsername_received_requests)); err != nil {
		fmt.Print("sentUsername requests")
		return errors.New("Failed to get sentUsername requests")
	}

	username_sent_requests = append(username_sent_requests, sentUsername)
	sentUsername_received_requests = append(sentUsername_received_requests, username)

	save_username_stmt := "UPDATE users SET sent_invitations = $1 WHERE username = $2"
	save_SentUsername_stmt := "UPDATE users SET received_invitations = $1 WHERE username = $2"

	_, err = db.Exec(save_username_stmt, pq.Array(username_sent_requests), username)

	if err != nil {
		fmt.Print("username")
		panic(err)
		return err
	}

	_, err = db.Exec(save_SentUsername_stmt, pq.Array(sentUsername_received_requests), sentUsername)

	if err != nil {
		fmt.Print("sentUsername")
		panic(err)
		return err
	}

	return err
}

func RemoveUserInvitation(username string, sentUsername string, token string) error {

	err := auth.CheckJWT(token, &username)

	if err != nil {
		return errors.New("invalid token")
	}

	db, err := configs.GetDB()
	if err != nil {
		fmt.Print("DB ERRRO")
		err = errors.New("DB connection error")
		return err
	}
	defer db.Close()

	var username_sent_requests []string
	var username_received_requests []string
	var sentUsername_sent_requests []string
	var sentUsername_received_requests []string

	username_stmt := "SELECT sent_invitations, received_invitations from users WHERE username=$1"
	sentUsername_stmt := "SELECT sent_invitations, received_invitations from users WHERE username=$1"

	if err := db.QueryRow(username_stmt, username).Scan(pq.Array(&username_sent_requests), pq.Array(&username_received_requests)); err != nil {
		return errors.New("Failed to get username requests")
	}
	if err := db.QueryRow(sentUsername_stmt, sentUsername).Scan(pq.Array(&sentUsername_sent_requests), pq.Array(&sentUsername_received_requests)); err != nil {
		return errors.New("Failed to get username requests")
	}

	if middleware.Contains(username_sent_requests, sentUsername) {
		username_sent_requests = middleware.Remove(username_sent_requests, sentUsername)
	}
	if middleware.Contains(username_received_requests, sentUsername) {
		username_received_requests = middleware.Remove(username_received_requests, sentUsername)
	}
	if middleware.Contains(sentUsername_sent_requests, username) {
		sentUsername_sent_requests = middleware.Remove(sentUsername_sent_requests, username)
	}
	if middleware.Contains(sentUsername_received_requests, username) {
		sentUsername_received_requests = middleware.Remove(sentUsername_received_requests, username)
	}

	save_username_stmt := "UPDATE users SET sent_invitations = $1, received_invitations = $2 WHERE username = $3"
	save_SentUsername_stmt := "UPDATE users SET sent_invitations = $1, received_invitations = $2 WHERE username = $3"

	_, err = db.Exec(save_username_stmt, pq.Array(username_sent_requests), pq.Array(username_received_requests), username)

	if err != nil {

		return err
	}

	_, err = db.Exec(save_SentUsername_stmt, pq.Array(sentUsername_sent_requests), pq.Array(sentUsername_received_requests), sentUsername)

	if err != nil {
		return err
	}

	return nil
}

func GetRecievedInvitationsByUsername(username string) (r_inv []string, e error) {
	db, err := configs.GetDB()
	if err != nil {
		err = errors.New("DB connection error")
		return nil, err
	}
	defer db.Close()

	var received_invitations []string

	stmt := "SELECT received_invitations FROM users WHERE username=$1"

	if err := db.QueryRow(stmt, username).Scan(pq.Array(&received_invitations)); err != nil {
		return nil, err
	}

	return received_invitations, err
}

func GetAllSentInvitationsHelper(username string) (r_inv []string, e error) {
	db, err := configs.GetDB()
	if err != nil {
		err = errors.New("DB connection error")
		return nil, err
	}
	defer db.Close()

	var sent_invitations []string

	stmt := "SELECT sent_invitations FROM users WHERE username=$1"

	if err := db.QueryRow(stmt, username).Scan(pq.Array(&sent_invitations)); err != nil {
		return nil, err
	}

	return sent_invitations, err
}

func GetAllSentInvitations(c *gin.Context) {

	var username string = c.Param("username")
	received_invitations, err := GetAllSentInvitationsHelper(username)

	if err != nil {
		c.JSON(400, gin.H{"error": err})
		return
	}

	c.JSON(200, gin.H{"data": received_invitations})
}

func GetAllReceivedInvitations(c *gin.Context) {
	var username string = c.Param("username")
	sent_invitations, err := GetRecievedInvitationsByUsername(username)

	if err != nil {
		c.JSON(400, gin.H{"error": err})
		return
	}

	c.JSON(200, gin.H{"data": sent_invitations})
}

func GetAllInvitations(c *gin.Context) {
	var username string = c.Param("username")
	sent_invitations, err := GetRecievedInvitationsByUsername(username)

	if err != nil {
		c.JSON(400, gin.H{"error": err})
		return
	}

	received_invitations, err := GetAllSentInvitationsHelper(username)

	if err != nil {
		c.JSON(400, gin.H{"error": err})
		return
	}

	c.JSON(200, gin.H{"sent": sent_invitations, "received": received_invitations})
}

func SendInvitation(c *gin.Context) {

	token := c.Param("token")
	username := c.Param("username")
	sentUsername := c.Param("sentUsername")

	err := SendUserInvitation(username, sentUsername, token)

	if err != nil {
		c.JSON(401, gin.H{"error": err})
		return
	}

	c.JSON(200, gin.H{"message": "request sent successfully"})

}

func DeleteInvitation(c *gin.Context) {
	token := c.Param("token")
	username := c.Param("username")
	friendUsername := c.Param("friendUsername")

	err := RemoveUserInvitation(username, friendUsername, token)

	if err != nil {
		c.JSON(400, gin.H{"error": err})
		return
	}

	c.JSON(200, gin.H{"message": "friend removed successfully"})
}
