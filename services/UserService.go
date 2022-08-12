package services

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/badoux/checkmail"
	"github.com/callummclu/Gocial-Media-Platform/configs"
	"github.com/callummclu/Gocial-Media-Platform/models"
	"github.com/gin-gonic/gin"
)

func GetAllUsers(c *gin.Context) {
	var users []models.ReturnedUser
	db, err := configs.GetDB()
	if err != nil {
		err = errors.New("DB connection error")
		c.JSON(400, gin.H{"error": err})
	}
	var query string = c.Query("searchParams")

	limit, err := strconv.Atoi(c.Query("itemsPerPage"))

	if err != nil {
		limit = 20
	}

	offset, err := strconv.Atoi(c.Query("page"))

	if err != nil {
		offset = 1
	}

	if offset < 1 {
		offset = 1
	}

	offset = (offset - 1) * limit

	rows, err := db.Query("select username, name, surname from users where strpos(username, $1) > 0 OR strpos(name, $1) > 0 OR strpos(surname, $1) > 0 LIMIT $2 OFFSET $3", query, limit, offset)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	}

	defer rows.Close()

	for rows.Next() {
		var (
			Username string
			Name     string
			Surname  string
		)

		if err := rows.Scan(&Username, &Name, &Surname); err != nil {
			fmt.Print(err)
		}

		users = append(users, models.ReturnedUser{
			Username: Username,
			Surname:  Surname,
			Name:     Name,
		})
	}

	if err != nil {
		c.JSON(400, gin.H{"error": "Cannot get users"})
		return
	}

	c.JSON(200, gin.H{"data": users})
}

func GetUserByUsername(c *gin.Context) {
	var user models.ReturnedUser

	var username string = c.Param("username")
	err := user.GetUserByUsername(username)

	if err != nil {
		c.JSON(400, gin.H{"error": "Cannot find user"})
		return
	}

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
