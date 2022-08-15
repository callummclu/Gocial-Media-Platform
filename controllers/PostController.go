package controllers

import (
	"github.com/callummclu/Gocial-Media-Platform/middleware"
	"github.com/callummclu/Gocial-Media-Platform/services"
)

func PostController() {
	api := Router.Group("post")
	{
		api.Use(middleware.CORSMiddleware("*"))

		api.GET("", services.GetPosts)
		api.GET(":username", services.GetPostByUsername)
		api.POST(":token", services.CreateNewPost)
		api.DELETE(":id/:token", services.DeleteOnePost)
		api.PUT(":id/:token", services.EditOnePost)

	}

}
