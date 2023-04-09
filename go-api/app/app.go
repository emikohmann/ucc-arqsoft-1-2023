package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-api/controllers"
)

const port = ":8080"

func StartApp() {
	router := gin.Default()
	// Then we will move this to url_mapping in order to organize the code
	router.GET("/items/:itemID", controllers.GetItem)
	err := router.Run(port)
	if err != nil {
		// Then we will move this to url_mapping in order to organize the code
		fmt.Println("error running app", err)
		return
	}
}
