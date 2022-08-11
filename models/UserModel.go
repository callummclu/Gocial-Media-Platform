package models

type User struct {
	Username string //UNIQUE
	Email    string //UNIQUE
	Password string
}
