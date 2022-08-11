package controllers

import "github.com/callummclu/Gocial-Media-Platform/services"

func ImageController() {
	api := Router.Group("image")
	{
		api.GET(":uuid", services.GetImageByUUID)
		api.POST("", services.SaveNewImage)
	}
}
