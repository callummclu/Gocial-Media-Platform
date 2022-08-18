package models

import (
	"errors"

	"github.com/callummclu/Gocial-Media-Platform/auth"
	"github.com/callummclu/Gocial-Media-Platform/configs"
)

type Image struct {
	uuid string
	uri  string
}

func NewImage() *Image {
	return new(Image)
}

func (i *Image) SaveImage(token string) error {

	var username string

	err := auth.CheckJWT(token, &username)

	if err != nil {
		err = errors.New("invalid user")
		return err
	}

	db, err := configs.GetDB()
	if err != nil {
		err = errors.New("DB connection error")
		return err
	}
	defer db.Close()

	insert_stmt, err := db.Prepare("INSERT INTO images {NEEDS TO BE FILLED IN} VALUES ({NEEDS TO BE FILLED IN})")

	if err != nil {
		return err
	}
	defer insert_stmt.Close()

	_, err = insert_stmt.Exec( /* NEEDS TO BE FILLED IN */ )

	return err
}
