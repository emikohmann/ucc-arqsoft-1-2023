package controllers

import (
	"fmt"
<<<<<<< HEAD
	"go-api/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetItem(ctx *gin.Context) {

	idString := ctx.Param("itemID")

	strconv.ParseInt(idString, 10, 64)
	if err != nil {
		fmt.Println("Error parsing item ID", err)
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	item := services.GetItem(id)
=======
	"github.com/gin-gonic/gin"
	"go-api/services"
	"net/http"
	"strconv"
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
>>>>>>> c0642656318807112e40c33dd6dd105d40bcd27c
	ctx.JSON(http.StatusOK, item)
}
