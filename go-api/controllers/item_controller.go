package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-api/services"
	"net/http"
	"strconv"
)

const (
	paramItemID = "itemID"
)

func GetItem(ctx *gin.Context) {
	// Get id param from URL as string
	idString := ctx.Param(paramItemID)

	// Convert string ID to int ID
	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		fmt.Println("Error parsing item ID", err)
		ctx.JSON(http.StatusBadRequest, err) // Then we will create a custom error struct
		return
	}

	// Call the service with int ID
	services.ItemClient = services.MlClient{}
	item, err := services.GetItem(id)
	if err != nil {
		fmt.Println("Error getting item", err)
		ctx.JSON(http.StatusInternalServerError, err) // Then we will create a custom error struct
		return
	}

	// Successful case
	ctx.JSON(http.StatusOK, item)
}