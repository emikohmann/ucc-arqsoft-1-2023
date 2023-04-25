package app

import (
	"fmt"
	"go-api/controllers"

	"github.com/gin-gonic/gin"
)

func StartApp() {

	router := gin.Default()
	router.GET("/items/:itemID", controllers.GetItem)
	err := router.Run(":8080")
	if err != nil {
		fmt.Println("Error running app", err)
		return
	}
}

const (
	port = ":8080"
)

/*
func StartApp() {
	router := gin.Default()

	// Then we will move this to url_mapping in order to organize the code
	router.GET("/items/:itemID", controllers.GetItem)

	err := router.Run(port)
	if err != nil {
		// Then we will replace this error with a panic since it's a non-recoverable error
		fmt.Println("Error running app", err)
		return
	}
}
*/
