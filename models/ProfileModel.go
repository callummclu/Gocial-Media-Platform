package models

import (
	"errors"

	"github.com/callummclu/Gocial-Media-Platform/configs"
)

type Profile struct {
	ID       int64  `json:"-";sql:"type:string REFERENCES users(username)"`
	Username string `json:"username"`
	// NEEDS
	// Posts []Post `json:"posts"`
	// Description string `json:"description"`
	// Name string `json:"name"`
	// Freinds []User `json:"friends"`

	// DISPLAYED AS
	// firstname 	Lastname 	@username	FRIENDS	FRIENDS
	// description 	description description	FRIENDS	FRIENDS
	// ------------------------------------	FRIENDS	FRIENDS
	// POSTS 		POSTS 		POSTS		FRIENDS	FRIENDS
	// POST			POSTS		POSTS		FRIENDS	FRIENDS
	// POST			POSTS		POSTS		FRIENDS	FRIENDS

}

func (p *Profile) GetProfileByUsername(query string) error {
	db, err := configs.GetDB()
	if err != nil {
		err = errors.New("DB connection error")
		return err
	}
	defer db.Close()
	err = db.QueryRow("SELECT * FROM profiles WHERE username = $1", query).Scan(&p.ID, &p.Username)

	return err
}

func NewProfile() *Profile {
	return new(Profile)
}

func (p *Profile) SaveProfile(username string) error {
	db, err := configs.GetDB()
	if err != nil {
		err = errors.New("DB connection error")
		return err
	}
	defer db.Close()

	insert_stmt, err := db.Prepare("INSERT INTO profiles (username) VALUES ($1)")

	if err != nil {
		return err
	}
	defer insert_stmt.Close()
	_, err = insert_stmt.Exec(username)
	return err
}
