package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
<<<<<<< HEAD
	"go-api/controllers"
=======
)

const (
	port = ":8080"
>>>>>>> master
)

func StartApp() {
	router := gin.Default()
<<<<<<< HEAD
	router.GET("/items/:itemID", controllers.GetItem)
	err := router.Run(":8080")
	if err != nil {
		fmt.Println("Error running app", err)
		return

	}

=======

	mapRoutes(router)

	err := router.Run(port)
	if err != nil {
		panic(fmt.Errorf("error running app: %w", err))
		return
	}
>>>>>>> master
}
