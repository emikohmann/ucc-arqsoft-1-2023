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

	strconv.ParseInt(idString, 10, 64)
	if err != nil {
		fmt.Println("Error parsing item ID", err)
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	item := services.GetItem(id)
	ctx.JSON(http.StatusOK, item)
}
