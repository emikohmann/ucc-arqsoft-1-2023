package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func StartApp() {
	router := gin.Default()
	err := router.Run(":8080")
	if err != nil {
		fmt.Println("Error running app", err)
		return

	}

}
