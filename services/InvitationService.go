package services

import (
	"github.com/callummclu/Gocial-Media-Platform/models"
	"github.com/gin-gonic/gin"
)

func GetAllSentInvitations(c *gin.Context) {

	var username string = c.Param("username")
	received_invitations, err := models.GetAllSentInvitations(username)

	if err != nil {
		c.JSON(400, gin.H{"error": err})
		return
	}

	c.JSON(200, gin.H{"data": received_invitations})
}

func GetAllReceivedInvitations(c *gin.Context) {
	var username string = c.Param("username")
	sent_invitations, err := models.GetRecievedInvitationsByUsername(username)

	if err != nil {
		c.JSON(400, gin.H{"error": err})
		return
	}

	c.JSON(200, gin.H{"data": sent_invitations})
}

func GetAllInvitations(c *gin.Context) {
	var username string = c.Param("username")
	sent_invitations, err := models.GetRecievedInvitationsByUsername(username)

	if err != nil {
		c.JSON(400, gin.H{"error": err})
		return
	}

	received_invitations, err := models.GetAllSentInvitations(username)

	if err != nil {
		c.JSON(400, gin.H{"error": err})
		return
	}

	c.JSON(200, gin.H{"sent": sent_invitations, "received": received_invitations})
}

func SendInvitation(c *gin.Context) {

	token := c.Param("token")
	username := c.Param("username")
	sentUsername := c.Param("sentUsername")

	err := models.SendUserInvitation(username, sentUsername, token)

	if err != nil {
		c.JSON(400, gin.H{"error": err})
		return
	}

	c.JSON(200, gin.H{"message": "request sent successfully"})

}

func DeleteInvitation(c *gin.Context) {
	token := c.Param("token")
	username := c.Param("username")
	friendUsername := c.Param("friendUsername")

	err := models.RemoveUserInvitation(username, friendUsername, token)

	if err != nil {
		c.JSON(400, gin.H{"error": err})
		return
	}

	c.JSON(200, gin.H{"message": "friend removed successfully"})
}
