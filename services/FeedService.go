package services

import (
	"github.com/callummclu/Gocial-Media-Platform/configs"
	"github.com/callummclu/Gocial-Media-Platform/models"
	"github.com/gin-gonic/gin"
)

func GetFriendsFeed(c *gin.Context) {
	db, err := configs.GetDB()
	if err != nil {
		c.JSON(504, gin.H{"error": "database error"})

	}
	defer db.Close()

	// ------------------------------------------------------------------------------------------
	/*
		GRAB USERS FRIENDS FROM DB

		username := c.Param("username")

		rows_friends, err := db.Query("SELECT friends from users WHERE username = $1")

		if err != nil {
			throw some error
		}

		defer rows_friends.Close()

		var friends []string

		if err := rows_friends.Scan(); err != nil {
			throw some error
		}

		users_list := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(friends)), ", "), "[]")

	*/

	users_list := ""

	// ------------------------------------------------------------------------------------------

	rows, err := db.Query("SELECT * from  posts where username IN ($1)", users_list)

	if err != nil {
		c.JSON(504, gin.H{"error": "database error"})
	}

	defer rows.Close()

	var posts []models.Post

	for rows.Next() {
		var (
			Username  string
			Title     string
			Content   string
			CreatedAt string
			Id        int64
		)

		if err := rows.Scan(&Username, &Title, &Content, &CreatedAt, &Id); err != nil {
			c.JSON(504, gin.H{"error": "database error"})
		}

		posts = append(posts, models.Post{
			Username:  Username,
			Title:     Title,
			Content:   Content,
			CreatedAt: CreatedAt,
			ID:        Id,
		})
	}
}
