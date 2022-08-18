package controllers

import (
	"github.com/callummclu/Gocial-Media-Platform/middleware"
	"github.com/callummclu/Gocial-Media-Platform/services"
)

func FeedController() {
	api := Router.Group("feed")
	{
		api.Use(middleware.CORSMiddleware("*"))

		api.GET("friends/:username", services.GetFriendsFeed)
	}
}
