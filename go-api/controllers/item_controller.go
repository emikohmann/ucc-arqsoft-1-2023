package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-api/services"
	"net/http"
	"strconv"
)

const paramItemId = "itemID"

func GetItem(ctx *gin.Context) {
	idString := ctx.Param(paramItemId)
	id, err := strconv.ParseInt(idString, 10, 64)

	if err != nil {
		fmt.Println("ERROR PASSING ITEM ID", err)
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	item, err := services.GetItem(id)
	if err != nil {
		fmt.Println("error getting items", err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, item)
}
