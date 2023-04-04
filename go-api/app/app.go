package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-api/controllers"
)

const port = ":8080"

func StartApp() {
	router := gin.Default()
	router.GET("/items/:itemID", controllers.GetItem)
	err := router.Run(port)
	if err != nil {
		fmt.Println("error running app", err)
		return
	}
}
