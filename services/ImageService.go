package services

import (
	"github.com/callummclu/Gocial-Media-Platform/configs"
	"github.com/gin-gonic/gin"
)

/*

	ALL FUNCTIONS HERE ARE NOT FINISHED
	AND STILL NEED TO BE

*/

func GetImageByUUID(c *gin.Context) {
	db, err := configs.GetDB()

	if err != nil {
		c.JSON(400, gin.H{"error": "DB Failed to Connect"})
		return
	}

	defer db.Close()
}

func SaveNewImage(c *gin.Context) {
	// SAVE UPLOADED IMAGE TO CDN.
	// SAVE URL TO CDN ALONGSIDE UUID TO IMAGE MODEL
	// RETURN THIS MODEL TO BE USED ELSEWHERE
}

func DeleteOneImage(c *gin.Context) {

}
