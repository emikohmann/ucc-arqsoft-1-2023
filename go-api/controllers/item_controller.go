package controllers

import (
	"fmt"
	"go-api/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

const paramItemID = "itemID"

func GetItem(ctx *gin.Context) {
	// Get id param from URL as string
	idString := ctx.Param(paramItemID)

	// Convert string ID to int ID
	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		fmt.Println("ERROR PASSING ITEM ID", err)
		ctx.JSON(http.StatusBadRequest, err) // then we weill create a custom error struct
		return
	}
	//call the services with int ID
	services.ItemClient = services.MlClient{}
	item, err := services.GetItem(id)
	if err != nil {
		fmt.Println("error getting items", err)
		ctx.JSON(http.StatusInternalServerError, err) // Then we will create a custom error struct
		return
	}
	//succesful case
	ctx.JSON(http.StatusOK, item)
}
