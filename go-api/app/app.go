package app

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func StartApp() {
	router := gin.Default()
	router.GET("/items/:itemID", controllers.GetItem())
	err := router.Run(":8080")  //bloquea la ejecucion de la aplicacion
	if err != nil {
		fmt.Println("error runig app", err)
		return
	}
}
