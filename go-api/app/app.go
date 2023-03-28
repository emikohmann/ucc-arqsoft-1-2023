package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-api/controllers"
)

func StartApp() {
	router := gin.Default()
	router.GET("/items/:itemID", controllers.GetItem)
	err := router.Run(":8080")
	if err != nil {
		fmt.Println("Error running app", err)
	}

}
