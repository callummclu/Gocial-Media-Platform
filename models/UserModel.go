package models

import (
	"database/sql"
	"errors"

	"github.com/callummclu/Gocial-Media-Platform/auth"
	"github.com/callummclu/Gocial-Media-Platform/configs"
)

type User struct {
	ID              int64  `json:"-"`
	Name            string `json:"name"`
	Surname         string `json:"surname"`
	Username        string `json:"username"`
	Email           string `json:"email"`
	EmailVerifiedAt string `json:"email_verified_at,omitempty"`
	Password        string `json:"password,omitempty"`
	CreatedAt       string `json:"created_at,omitempty"`
	UpdatedAt       string `json:"updated_at,omitempty"`
}

type LogInUser struct {
	User     string `json:"username"`
	Password string `json:"password"`
}

type ReturnedUser struct {
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Username string `json:"username"`
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

			profile := NewProfile()

			err = profile.SaveProfile(u.Username)

			if err != nil {
				return err
			}

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

	err = db.QueryRow("SELECT username, name, surname FROM users WHERE username = $1", query).Scan(&u.Name, &u.Surname, &u.Username)

	return err
}
