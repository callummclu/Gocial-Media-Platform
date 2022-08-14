package controllers

import (
	"github.com/callummclu/Gocial-Media-Platform/middleware"
	"github.com/callummclu/Gocial-Media-Platform/services"
)

func ProfileController() {
	api := Router.Group("profile")
	{
		api.Use(middleware.CORSMiddleware("*"))

		api.GET("", services.GetProfileByUsername)
	}
}
