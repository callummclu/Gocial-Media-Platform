package models

import (
	"errors"

	"github.com/callummclu/Gocial-Media-Platform/auth"
	"github.com/callummclu/Gocial-Media-Platform/configs"
)

type Post struct {
	ID        int64  `json:"id";sql:"type:string REFERENCES users(username)"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Username  string `json:"username"`
	CreatedAt string `json:"created_at,omitempty"`
}

func NewPost() *Post {
	return new(Post)
}

func (u *Post) SavePost(token string) error {

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

	insert_stmt, err := db.Prepare("INSERT INTO posts (title,content,username) VALUES ($1,$2,$3)")

	if err != nil {
		return err
	}
	defer insert_stmt.Close()

	_, err = insert_stmt.Exec(u.Title, u.Content, username)

	return err
}
