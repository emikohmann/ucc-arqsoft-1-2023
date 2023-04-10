package app

import (
	"github.com/gin-gonic/gin"
	"go-api/controllers"
)

const (
	pathGetItem = "/items/:itemID"
)

func mapRoutes(router *gin.Engine) {
	router.GET(pathGetItem, controllers.GetItem)
}
