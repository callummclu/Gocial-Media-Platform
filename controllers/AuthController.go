package controllers

import (
	"github.com/callummclu/Gocial-Media-Platform/middleware"
	"github.com/callummclu/Gocial-Media-Platform/services"
)

func AuthController() {
	api := Router.Group("auth")
	{
		api.Use(middleware.CORSMiddleware("*"))
		api.POST("login", services.LoginUser)
		api.GET(":token", services.CheckUser)
		// Logout should be done client side by removing JWT
	}

}
