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

		api.GET("invitation/:username/sent", services.GetAllSentInvitations)
		api.GET("invitation/:username/received", services.GetAllReceivedInvitations)
		api.GET("invitation/:username", services.GetAllInvitations)

		api.POST("invitation/:username/:sentUsername/:token", services.SendInvitation)
		api.DELETE("invitation/:username/:sentUsername", services.DeleteInvitation)
		api.POST("friends/:username/:sentUsername/:token/accept", services.AcceptFriendRequest)

		api.GET("friends/:username", services.GetUsersFriends)
		api.DELETE("friends/:username/:friendUsername", services.DeleteFriend)

	}
}
