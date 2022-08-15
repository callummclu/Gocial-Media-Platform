package controllers

import (
	"github.com/callummclu/Gocial-Media-Platform/configs"
	"github.com/gin-gonic/gin"
)

var Router *gin.Engine = gin.Default()

func BaseController() {
	UserController()
	AuthController()
	ImageController()
	Router.Run(configs.EnvPORT())
}
