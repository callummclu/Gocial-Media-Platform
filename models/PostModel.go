package models

import (
	"errors"
	"fmt"

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

func GetPostsByUsername(query string) (posts []Post, e error) {
	db, err := configs.GetDB()
	if err != nil {
		err = errors.New("DB connection error")
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT username, title, content, created_at, id FROM posts WHERE username = $1 ORDER BY created_at DESC", query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var p []Post

	for rows.Next() {
		var (
			Username  string
			Title     string
			Content   string
			CreatedAt string
			Id        int64
		)

		if err := rows.Scan(&Username, &Title, &Content, &CreatedAt, &Id); err != nil {
			fmt.Print(err)
		}

		p = append(p, Post{
			Username:  Username,
			Title:     Title,
			Content:   Content,
			CreatedAt: CreatedAt,
			ID:        Id,
		})
	}

	return p, err
}

func DeletePostById(id int64, token string, username string) error {

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

	_, err = db.Query("DELETE FROM posts WHERE id = $1", id)

	return err
}
