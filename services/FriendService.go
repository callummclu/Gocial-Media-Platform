package services

import (
	"github.com/callummclu/Gocial-Media-Platform/models"
	"github.com/gin-gonic/gin"
)

func AcceptFriendRequest(c *gin.Context) {

	token := c.Param("token")
	username := c.Param("username")
	friendUsername := c.Param("friendUsername")

	err := models.AcceptUserFriendRequest(username, friendUsername, token)

	if err != nil {
		c.JSON(400, gin.H{"error": err})
		return
	}

	c.JSON(200, gin.H{"message": "friend added successfully"})
}

func GetUsersFriends(c *gin.Context) {
	var username string = c.Param("username")
	friends, err := models.GetAllFriends(username)

	if err != nil {
		c.JSON(400, gin.H{"error": err})
		return
	}

	c.JSON(200, gin.H{"friends": friends})

}

func DeleteFriend(c *gin.Context) {
	token := c.Param("token")
	username := c.Param("username")
	friendUsername := c.Param("friendUsername")

	err := models.RemoveUserFriend(username, friendUsername, token)

	if err != nil {
		c.JSON(400, gin.H{"error": err})
		return
	}

	c.JSON(200, gin.H{"message": "friend removed successfully"})
}
