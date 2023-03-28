package app

import "github.com/gin-gonic/gin"
import "fmt"

func StartApp() {
	router := gin.Default()
	err := router.Run(":8080")
	if err != nil {
		fmt.Println("Error ", err)
		return
	}

}
