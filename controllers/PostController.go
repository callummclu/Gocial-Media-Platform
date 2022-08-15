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
		api.POST("", services.CreateNewPost)
		api.DELETE(":id", services.DeleteOnePost)
		api.PUT(":id", services.EditOnePost)

	}

}
