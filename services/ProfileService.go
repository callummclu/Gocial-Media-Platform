package services

import (
	"fmt"

	"github.com/callummclu/Gocial-Media-Platform/models"
	"github.com/gin-gonic/gin"
)

func GetProfileByUsername(c *gin.Context) {
	var profile models.Profile

	var username string = c.Param("username")
	err := profile.GetProfileByUsername(username)

	if err != nil {
		fmt.Println(err)
		fmt.Println(profile)
		c.JSON(400, gin.H{"error": "Cannot find profile"})
		return
	}

	c.JSON(200, gin.H{"data": profile})
}
