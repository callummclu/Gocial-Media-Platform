package controllers

import (
	"github.com/callummclu/Gocial-Media-Platform/middleware"
	"github.com/callummclu/Gocial-Media-Platform/services"
)

func UserController() {
	api := Router.Group("user")
	{
		api.Use(middleware.CORSMiddleware("*"))

		api.GET("", services.GetAllUsers)
		api.GET(":username", services.GetUserByUsername)
		api.POST("", services.CreateNewUser)

		// NEEDS AUTH MIDDLEWARE
		api.DELETE("", services.DeleteOneUser)

		// NEEDS AUTH MIDDLEWARE
		api.PUT(":id", services.EditOneUser)
	}
}
