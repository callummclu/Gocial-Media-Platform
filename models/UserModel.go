package models

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/callummclu/Gocial-Media-Platform/auth"
	"github.com/callummclu/Gocial-Media-Platform/configs"
	"github.com/callummclu/Gocial-Media-Platform/middleware"
	"github.com/lib/pq"
)

type User struct {
	ID                  int64    `json:"-"`
	Name                string   `json:"name"`
	Surname             string   `json:"surname"`
	Username            string   `json:"username"`
	Email               string   `json:"email"`
	Description         string   `json:"description"`
	DisplayImage        string   `json:"display_image"`
	Friends             []string `json:"friends"`
	ReceivedInvitations []string `json:"received_invitations"`
	SentInvitations     []string `json:"sent_invitations"`
	EmailVerifiedAt     string   `json:"email_verified_at,omitempty"`
	Password            string   `json:"password,omitempty"`
	CreatedAt           string   `json:"created_at,omitempty"`
}

type LogInUser struct {
	User     string `json:"username"`
	Password string `json:"password"`
}

type ReturnedUser struct {
	Name                string   `json:"name"`
	Surname             string   `json:"surname"`
	Username            string   `json:"username"`
	DisplayImage        string   `json:"display_image"`
	Description         string   `json:"description"`
	Friends             []string `json:"friends"`
	ReceivedInvitations []string `json:"received_invitations"`
}

type ReturnedUsers struct {
	Users []ReturnedUser `json:"users"`
}

func NewUser() *User {
	return new(User)
}

func NewLogInUser() *LogInUser {
	return new(LogInUser)
}

func (u *User) SaveUser() error {
	db, err := configs.GetDB()
	if err != nil {
		err = errors.New("DB connection error")
		return err
	}
	defer db.Close()
	/*
	 Check if the user already exists
	*/
	var (
		username string
		email    string
	)
	stmt, err := db.Prepare("SELECT username,email FROM users WHERE username = $1 OR email = $2")
	if err != nil {
		return err
	}
	defer stmt.Close()
	err = stmt.QueryRow(u.Username, u.Email).Scan(&username, &email)
	if err != nil {
		if err == sql.ErrNoRows {
			hashedPassword, err := auth.HashPassword(u.Password)
			if err != nil {
				return err
			}
			//Add the new user
			insert_stmt, err := db.Prepare("INSERT INTO users (name,surname,username,email,password,display_image,description,friends,received_invitations) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$8)")

			if err != nil {
				return err
			}
			defer insert_stmt.Close()
			var emptyString []string = nil
			_, err = insert_stmt.Exec(u.Name, u.Surname, u.Username, u.Email, hashedPassword, "", "", pq.Array(emptyString))

			return err
		} else {
			return err
		}
	} else {
		err = errors.New("User already exists")
		return err
	}
	return err
}

func (u *LogInUser) UserLogin() (string, error) {
	db, err := configs.GetDB()
	if err != nil {
		err = errors.New("DB connection error")
		return "", err
	}
	defer db.Close()
	stmt, err := db.Prepare("SELECT username,password FROM users WHERE username = $1 OR email = $1")
	if err != nil {
		return "", err
	}
	defer stmt.Close()
	var (
		username string
		password string
	)
	err = stmt.QueryRow(u.User).Scan(&username, &password)

	if err != nil {
		return "", err
	}
	err = auth.CheckPassword(password, u.Password)

	if err != nil {
		err = errors.New("Wrong password")
		return "", err
	}
	return username, err
}

func (u *ReturnedUser) GetUserByUsernameQuery(query string) error {
	db, err := configs.GetDB()
	if err != nil {
		err = errors.New("DB connection error")
		return err
	}
	defer db.Close()

	err = db.QueryRow("SELECT username, name, surname, display_image, description, friends, received_invitations FROM users WHERE username = $1", query).Scan(&u.Name, &u.Surname, &u.Username, &u.DisplayImage, &u.Description, pq.Array(&u.Friends), pq.Array(&u.ReceivedInvitations))

	return err
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

func GetAllSentInvitations(username string) (r_inv []string, e error) {
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

	middleware.Remove(username_friends, sentUsername)
	middleware.Remove(sentUsername_friends, username)

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
		fmt.Print("username requests")
		return errors.New("Failed to get username requests")
	}
	if err := db.QueryRow(sentUsername_stmt, sentUsername).Scan(pq.Array(&sentUsername_sent_requests), pq.Array(&sentUsername_received_requests)); err != nil {
		fmt.Print("username requests")
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

	return nil
}
