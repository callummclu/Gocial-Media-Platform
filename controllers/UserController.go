package controllers

import "github.com/callummclu/Gocial-Media-Platform/services"

func UserController() {
	api := Router.Group("user")
	{
		//
		api.GET("", services.GetAllUsers)

		//
		api.GET(":id", services.GetUserByUsername)

		// TAKES USER MODEL in body
		api.POST("", services.CreateNewUser)

		//
		api.DELETE("", services.DeleteOneUser)

		//
		api.PUT(":id", services.EditOneUser)
	}
}
