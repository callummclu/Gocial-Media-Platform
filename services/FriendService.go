package services

import (
	"github.com/callummclu/Gocial-Media-Platform/models"
	"github.com/gin-gonic/gin"
)

func AcceptFriendRequest(c *gin.Context) {

	token := c.Param("token")
	username := c.Param("username")
	sentUsername := c.Param("sentUsername")

	err := models.AcceptUserFriendRequest(username, sentUsername, token)

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
	// Check token is valid

	// GET friends from username
	// GET friends from sentUsername

	// friends_username = remove sentUsername
	// friends_sentUsername = remove username

	// UPDATE users SET friends = friends_username WHERE username=username
	// UPDATE users SET friends = friends_sentUsername WHERE username=sentUsername
}
