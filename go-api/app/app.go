package app

import (
	"fmt"
<<<<<<< HEAD

	"github.com/gin-gonic/gin"
)

func StartApp() {
	router := gin.Default()
	router.GET("/items/:itemID", controllers.GetItem())
	err := router.Run(":8080")  //bloquea la ejecucion de la aplicacion
	if err != nil {
		fmt.Println("error runig app", err)
=======
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
>>>>>>> c0642656318807112e40c33dd6dd105d40bcd27c
		return
	}
}
