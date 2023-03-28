package controllers

import (
	"fmt"
	"go-api/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetItem(ctx *gin.Context) {
	idString := ctx.Param("itemID")
	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		fmt.Println("Error parasing item id", err)
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	item, err := services.GetItem(id)
	if err != nil {
		fmt.Println("Error getting item", err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, item)
}
