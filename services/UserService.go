package services

import (
	"strings"

	"github.com/badoux/checkmail"
	"github.com/callummclu/Gocial-Media-Platform/models"
	"github.com/gin-gonic/gin"
)

func GetAllUsers(c *gin.Context) {
	var users []models.ReturnedUser

	// takes query params
	// THE SQL SHOULD SEARCH FOR
	// (
	// 	SELECT username, name, surname
	// 	FROM users
	// 	WHERE strpos(username, $1) > 0
	// 		OR strpos(name, $1) > 0
	// 		OR strpos(surname, $1) > 0
	// )
	// including searchPhrase
	// including page no
	// including items per page
	// i.e. domain.com/?searchPhrase=callum&page=1&itemsPerPage=20

	c.JSON(200, gin.H{"data": users})
}
func GetUserByUsername(c *gin.Context) {
	var user models.ReturnedUser

	// SELECT username, name, surname
	// FROM users
	// WHERE username = 'callummclu'

	c.JSON(200, gin.H{"data": user})

}

func CreateNewUser(c *gin.Context) {
	user := models.NewUser()
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	if strings.TrimSpace(user.Name) == "" && strings.TrimSpace(user.Surname) == "" && strings.TrimSpace(user.Username) == "" && strings.TrimSpace(user.Email) == "" && strings.TrimSpace(user.Password) == "" {
		c.JSON(400, gin.H{
			"error": "Fill in all the require fields",
		})
		return
	}
	user.Username = strings.TrimSpace(user.Username)
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	user.Password = strings.TrimSpace(user.Password)

	if checkmail.ValidateFormat(user.Email) != nil {
		c.JSON(400, gin.H{
			"error": "Invalid email format",
		})
		return
	}
	err := user.SaveUser()
	if err != nil {
		if err.Error() == "User already exists" {
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
	c.JSON(200, gin.H{
		"message": "User registered successfully",
	})
}
func DeleteOneUser(c *gin.Context) {}
func EditOneUser(c *gin.Context)   {}
