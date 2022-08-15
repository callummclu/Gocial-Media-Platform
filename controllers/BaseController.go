package controllers

import (
	"github.com/callummclu/Gocial-Media-Platform/configs"
	"github.com/gin-gonic/gin"
)

var Router *gin.Engine = gin.Default()

func BaseController() {
	// Router.Use(static.Serve("/", static.LocalFile("./client/build", true)))

	UserController()
	PostController()
	AuthController()
	ImageController()
	Router.Run(configs.EnvPORT())
}
