package controllers

import (
	"github.com/gin-gonic/gin"
)

var Router *gin.Engine = gin.Default()

func BaseController() {
	UserController()
}
