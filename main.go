package main

import (
	"github.com/callummclu/Gocial-Media-Platform/configs"
	"github.com/callummclu/Gocial-Media-Platform/controllers"
	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func main() {
	// configs.ConnectDB()
	controllers.BaseController()
	configs.RunServer(&Router)
}
