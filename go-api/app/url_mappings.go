package app

import (
	"go-api/controllers"

	"github.com/gin-gonic/gin"
)

const (
	pathGetItem = "/items/:itemID"
)

func mapRoutes(router *gin.Engine) {
	router.GET(pathGetItem, controllers.GetItem)
}
