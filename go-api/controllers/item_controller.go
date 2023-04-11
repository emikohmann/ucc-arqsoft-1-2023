package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-api/services"
	"go-api/services/clients"
	"net/http"
	"strconv"
)

const (
	paramItemID = "itemID"
)

func init() {
	// Here we define that for the real application we will use HTTPClient as MLClient
	services.MLClient = clients.HTTPClient{}
}

func GetItem(ctx *gin.Context) {
	// Get id param from URL as string
	idString := ctx.Param(paramItemID)

	// Convert string ID to int ID
	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Errorf("error parsing item ID: %w", err))
		return
	}

	// Call the service with int ID
	item, err := services.GetItem(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, fmt.Errorf("error getting item: %w", err))
		return
	}

	// Successful case
	ctx.JSON(http.StatusOK, item)
}
