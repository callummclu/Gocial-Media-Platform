package services

import (
	"strings"

	"github.com/callummclu/Gocial-Media-Platform/auth"
	"github.com/callummclu/Gocial-Media-Platform/models"
	"github.com/gin-gonic/gin"
)

func LoginUser(c *gin.Context) {
	user_login := models.NewLogInUser()
	if err := c.ShouldBindJSON(&user_login); err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	if strings.TrimSpace(user_login.User) == "" && strings.TrimSpace(user_login.Password) == "" {
		c.JSON(400, gin.H{
			"error": "Fill in all the require fields",
		})
		return
	}
	username, err := user_login.UserLogin()
	if err != nil {
		if err.Error() == "The user does not exist" || err.Error() == "Wrong password" {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	token, err := auth.GenerateJWT(username)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"token": token,
	})
}

func CheckUser(c *gin.Context) {
	token := c.Param("token")
	if strings.TrimSpace(token) == "" {
		c.JSON(400, gin.H{
			"error": "Token not provided",
		})
		return
	}
	var username string

	err := auth.CheckJWT(token, &username)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"isAuthenticated": true,
		"username":        username,
	})

	// needs to return specific user
}
