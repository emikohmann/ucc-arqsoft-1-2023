package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

const (
	port = ":8080"
)

func StartApp() {
	router := gin.Default()

	mapRoutes(router)

	err := router.Run(port)
	if err != nil {
		panic(fmt.Errorf("error running app: %w", err))
		return
	}
}
