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
		api.DELETE(":token", services.DeleteOneUser)
		// NEEDS AUTH MIDDLEWARE
		// api.PUT(":id/:token", services.EditOneUser)

		api.GET("invitation/:username/sent", services.GetAllSentInvitations)
		api.GET("invitation/:username/received", services.GetAllReceivedInvitations)
		api.GET("invitation/:username", services.GetAllInvitations)
		api.POST("invitation/:username/:sentUsername/:token", services.SendInvitation)
		api.DELETE("invitation/:username/:sentUsername/:token/remove", services.DeleteInvitation)

		api.GET("friends/:username", services.GetUsersFriends)
		api.POST("friends/:username/:friendUsername/:token/accept", services.AcceptFriendRequest)
		api.DELETE("friends/:username/:friendUsername/:token", services.DeleteFriend)
	}
}
