package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-api/services"
	"net/http"
	"strconv"
)

func GetItem(ctx *gin.Context) {

	idString := ctx.Param("itemID")
	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		fmt.Println(err.Error())
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	services.ItemClient = services.MlClient{}
	item, error := services.GetItem(id)
	if error != nil {
		fmt.Println("Error getting item", error)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, item)
}
