package app

import (
	"fmt"
	"go-api/controllers"
	"github.com/gin-gonic/gin"
)

func StartApp(){


	router:=gin.Default()
	router.GET("/items/:itemsID", controllers.GetItem)
	err:=router.Run(":8080")

	if err != nil{ //bloquea ejecucion
		fmt.Println("Error", err)
		return 
	}

}