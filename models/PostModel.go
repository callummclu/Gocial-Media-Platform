package models

import (
	"errors"
	"fmt"

	"github.com/callummclu/Gocial-Media-Platform/auth"
	"github.com/callummclu/Gocial-Media-Platform/configs"
)

type Post struct {
	ID       int64  `json:"-";sql:"type:string REFERENCES users(username)"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	Username string `json:"username"`
}

func NewPost() *Post {
	return new(Post)
}

func (u *Post) SavePost(token string) error {

	var username string

	err := auth.CheckJWT(token, &username)

	if err != nil {
		return errors.New("invalid user")
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

func GetPostsByUsername(query string) (posts []Post, e error) {
	db, err := configs.GetDB()
	if err != nil {
		err = errors.New("DB connection error")
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT username, title, content FROM posts WHERE username = $1", query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var p []Post

	for rows.Next() {
		var (
			Username string
			Title    string
			Content  string
		)

		if err := rows.Scan(&Username, &Title, &Content); err != nil {
			fmt.Print(err)
		}

		p = append(p, Post{
			Username: Username,
			Title:    Title,
			Content:  Content,
		})
	}

	return p, err
}
