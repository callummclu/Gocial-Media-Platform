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

func GetAllFriends(username string) (f_list []string, e error) {
	db, err := configs.GetDB()
	if err != nil {
		err = errors.New("DB connection error")
		return nil, err
	}
	defer db.Close()

	var friends []string

	stmt := "SELECT friends FROM users WHERE username=$1"

	if err := db.QueryRow(stmt, username).Scan(pq.Array(&friends)); err != nil {
		return nil, err
	}

	return friends, err
}

func RemoveUserFriend(username string, sentUsername string, token string) error {

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

	var username_friends []string
	var sentUsername_friends []string

	username_friends_stmt := "SELECT friends from users WHERE username=$1"
	sentUsername_friends_stmt := "SELECT friends from users WHERE username=$1"

	if err := db.QueryRow(username_friends_stmt, username).Scan(pq.Array(&username_friends)); err != nil {
		return errors.New("Failed to get username friends")
	}
	if err := db.QueryRow(sentUsername_friends_stmt, sentUsername).Scan(pq.Array(&sentUsername_friends)); err != nil {
		return errors.New("Failed to get username friends")
	}

	username_friends = middleware.Remove(username_friends, sentUsername)
	sentUsername_friends = middleware.Remove(sentUsername_friends, username)

	save_username_friends_stmt := "UPDATE users SET friends = $1 WHERE username = $2"
	save_SentUsername_friends_stmt := "UPDATE users SET friends = $1 WHERE username = $2"

	_, err = db.Exec(save_username_friends_stmt, pq.Array(username_friends), username)

	if err != nil {
		fmt.Print("username")
		panic(err)
		return err
	}

	_, err = db.Exec(save_SentUsername_friends_stmt, pq.Array(sentUsername_friends), sentUsername)

	if err != nil {
		fmt.Print("sentUsername")
		panic(err)
		return err
	}
	return nil
}

func AcceptUserFriendRequest(username string, sentUsername string, token string) error {

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

	var username_received_requests []string

	username_received_stmt := "SELECT received_invitations from users WHERE username=$1"

	if err := db.QueryRow(username_received_stmt, username).Scan(pq.Array(&username_received_requests)); err != nil {
		return errors.New("Failed to get username requests")
	}

	if !middleware.Contains(username_received_requests, sentUsername) {
		fmt.Println("user hasnt sent you a request")
		return errors.New("user hasnt sent you a request")
	}

	username_received_requests = middleware.Remove(username_received_requests, sentUsername)

	var username_sent_request []string
	var sentUsername_sent_requests []string
	var sentUsername_received_requests []string

	username_sent_stmt := "SELECT sent_invitations from users WHERE username=$1"
	sentUsername_sent_stmt := "SELECT sent_invitations from users WHERE username=$1"
	sentUsername_received_stmt := "SELECT received_invitations from users WHERE username=$1"

	if err := db.QueryRow(username_sent_stmt, username).Scan(pq.Array(&username_sent_request)); err != nil {
		return errors.New("Failed to get username requests")
	}
	if err := db.QueryRow(sentUsername_sent_stmt, sentUsername).Scan(pq.Array(&sentUsername_sent_requests)); err != nil {
		return errors.New("Failed to get username requests")
	}
	if err := db.QueryRow(sentUsername_received_stmt, sentUsername).Scan(pq.Array(&sentUsername_received_requests)); err != nil {
		return errors.New("Failed to get username requests")
	}

	if middleware.Contains(username_sent_request, sentUsername) {
		username_sent_request = middleware.Remove(username_sent_request, sentUsername)
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

	_, err = db.Exec(save_username_stmt, pq.Array(username_sent_request), pq.Array(username_received_requests), username)

	if err != nil {
		fmt.Print("username")
		panic(err)
		return err
	}

	_, err = db.Exec(save_SentUsername_stmt, pq.Array(sentUsername_sent_requests), pq.Array(sentUsername_received_requests), sentUsername)

	if err != nil {
		fmt.Print("sentUsername")
		panic(err)
		return err
	}

	var username_friends []string
	var sentUsername_friends []string

	username_friends_stmt := "SELECT friends from users WHERE username=$1"
	sentUsername_friends_stmt := "SELECT friends from users WHERE username=$1"

	if err := db.QueryRow(username_friends_stmt, username).Scan(pq.Array(&username_friends)); err != nil {
		return errors.New("Failed to get username friends")
	}
	if err := db.QueryRow(sentUsername_friends_stmt, sentUsername).Scan(pq.Array(&sentUsername_friends)); err != nil {
		return errors.New("Failed to get username friends")
	}

	username_friends = append(username_friends, sentUsername)
	sentUsername_friends = append(sentUsername_friends, username)

	save_username_friends_stmt := "UPDATE users SET friends = $1 WHERE username = $2"
	save_SentUsername_friends_stmt := "UPDATE users SET friends = $1 WHERE username = $2"

	_, err = db.Exec(save_username_friends_stmt, pq.Array(username_friends), username)

	if err != nil {
		return err
	}

	_, err = db.Exec(save_SentUsername_friends_stmt, pq.Array(sentUsername_friends), sentUsername)

	if err != nil {
		return err
	}

	return nil
}

func AcceptFriendRequest(c *gin.Context) {

	token := c.Param("token")
	username := c.Param("username")
	friendUsername := c.Param("friendUsername")

	err := AcceptUserFriendRequest(username, friendUsername, token)

	if err != nil {
		c.JSON(400, gin.H{"error": err})
		return
	}

	c.JSON(200, gin.H{"message": "friend added successfully"})
}

func GetUsersFriends(c *gin.Context) {
	var username string = c.Param("username")
	friends, err := GetAllFriends(username)

	if err != nil {
		c.JSON(400, gin.H{"error": err})
		return
	}

	c.JSON(200, gin.H{"friends": friends})
}

func DeleteFriend(c *gin.Context) {
	token := c.Param("token")
	username := c.Param("username")
	friendUsername := c.Param("friendUsername")

	err := RemoveUserFriend(username, friendUsername, token)

	if err != nil {
		c.JSON(400, gin.H{"error": err})
		return
	}

	c.JSON(200, gin.H{"message": "friend removed successfully"})
}
