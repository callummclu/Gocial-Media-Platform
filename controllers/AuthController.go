package controllers

import "github.com/callummclu/Gocial-Media-Platform/services"

func AuthController() {
	api := Router.Group("auth")
	{
		api.POST("login", services.LoginUser)
		api.POST("logout", services.Logout)
		api.GET(":token", services.CheckUser)
	}

}
