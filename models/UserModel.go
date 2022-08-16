package models

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/callummclu/Gocial-Media-Platform/auth"
	"github.com/callummclu/Gocial-Media-Platform/configs"
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
	Name         string `json:"name"`
	Surname      string `json:"surname"`
	Username     string `json:"username"`
	DisplayImage string `json:"display_image"`
	Description  string `json:"description"`
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
			insert_stmt, err := db.Prepare("INSERT INTO users (name,surname,username,email,password) VALUES ($1,$2,$3,$4,$5)")

			if err != nil {
				return err
			}
			defer insert_stmt.Close()
			_, err = insert_stmt.Exec(u.Name, u.Surname, u.Username, u.Email, hashedPassword)

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

func (u *ReturnedUser) GetUserByUsername(query string) error {
	db, err := configs.GetDB()
	if err != nil {
		err = errors.New("DB connection error")
		return err
	}
	defer db.Close()

	err = db.QueryRow("SELECT username, name, surname, display_image, description FROM users WHERE username = $1", query).Scan(&u.Name, &u.Surname, &u.Username, &u.DisplayImage, &u.Description)

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

func SendUserInvitation(username string, sentUsername string) error {
	db, err := configs.GetDB()
	if err != nil {
		fmt.Print("DB ERRRO")
		err = errors.New("DB connection error")
		return err
	}
	defer db.Close()

	var username_sent_requests []string
	var sentUsername_received_requests []string

	username_stmt := "SELECT sent_invitations from users WHERE username=$1"
	sentUsername_stmt := "SELECT received_invitations from users WHERE username=$1"

	if err := db.QueryRow(username_stmt, username).Scan(pq.Array(&username_sent_requests)); err != nil {
		fmt.Print("username requests")
		return errors.New("Failed to get username requests")

	}

	// Check to see if sentUsername is in username_sent_request
	// Return already sent

	/*
		NEED A CHECK TO SEE IF sentUsername is already in recieved_requests for username to ensure this data is not doubled up
	*/

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
