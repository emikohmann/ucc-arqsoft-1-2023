package app

import (
	"go-api/controllers"

	"github.com/gin-gonic/gin"
)

const (
	pathGetItem = "/items/:itemID"
	pathAuth    = "/auth"
)

func mapRoutes(router *gin.Engine) {
	router.GET(pathGetItem, controllers.GetItem)
	router.POST(pathAuth, controllers.Auth)
}
