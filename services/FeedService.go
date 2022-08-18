package services

import (
	"fmt"
	"strings"

	"github.com/callummclu/Gocial-Media-Platform/configs"
	"github.com/callummclu/Gocial-Media-Platform/models"
	"github.com/gin-gonic/gin"
)

func GetFriendsFeed(c *gin.Context) {
	db, err := configs.GetDB()
	if err != nil {
		c.JSON(504, gin.H{"error": "database error"})
		return

	}
	defer db.Close()

	username := c.Param("username")

	rows_friends, err := db.Query("SELECT friends from users WHERE username = $1", username)

	if err != nil {
		c.JSON(500, gin.H{"error": "db error", "details": err})
		return
	}

	defer rows_friends.Close()

	var friends []string

	for rows_friends.Next() {
		var (
			Friend string
		)

		if err := rows_friends.Scan(&Friend); err != nil {
			c.JSON(504, gin.H{"error": "database friends error"})
			return
		}

		friends = append(friends, Friend)
	}

	users_list := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(friends)), ", "), "")

	users_list = strings.ReplaceAll(users_list, "{", "")
	users_list = strings.ReplaceAll(users_list, "}", "")
	users_list = strings.ReplaceAll(users_list, "[", "")
	users_list = strings.ReplaceAll(users_list, "]", "")
	users_list = strings.ReplaceAll(users_list, ",", ",")

	rows, err := db.Query("SELECT username,title,content,created_at,id from posts where username = ANY(ARRAY[$1]) ORDER BY created_at DESC", users_list)

	if err != nil {
		c.JSON(504, gin.H{"error": "database posts error"})
		return
	}

	defer rows.Close()

	var p []models.Post

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

		p = append(p, models.Post{
			Username:  Username,
			Title:     Title,
			Content:   Content,
			CreatedAt: CreatedAt,
			ID:        Id,
		})
	}

	c.JSON(200, gin.H{"data": p})

}
