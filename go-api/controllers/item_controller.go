package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-api/services"
<<<<<<< HEAD
=======
	"go-api/services/clients"
>>>>>>> master
	"net/http"
	"strconv"
)

const (
	paramItemID = "itemID"
)

<<<<<<< HEAD
=======
func init() {
	// Here we define that for the real application we will use HTTPClient as MLClient
	services.MLClient = clients.HTTPClient{}
}

>>>>>>> master
func GetItem(ctx *gin.Context) {
	// Get id param from URL as string
	idString := ctx.Param(paramItemID)

	// Convert string ID to int ID
	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
<<<<<<< HEAD
		fmt.Println("Error parsing item ID", err)
		ctx.JSON(http.StatusBadRequest, err) // Then we will create a custom error struct
=======
		ctx.JSON(http.StatusBadRequest, fmt.Errorf("error parsing item ID: %w", err))
>>>>>>> master
		return
	}

	// Call the service with int ID
<<<<<<< HEAD
	services.ItemClient = services.MlClient{}
	item, err := services.GetItem(id)
	if err != nil {
		fmt.Println("Error getting item", err)
		ctx.JSON(http.StatusInternalServerError, err) // Then we will create a custom error struct
=======
	item, err := services.GetItem(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, fmt.Errorf("error getting item: %w", err))
>>>>>>> master
		return
	}

	// Successful case
	ctx.JSON(http.StatusOK, item)
}
