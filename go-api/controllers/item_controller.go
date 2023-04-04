package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-api/services"
	"strconv" 
	"net/http")

//contexto que no recibe nada
func GetItem(ctx *gin.Context){

	idString:= ctx.Param("itemsID")

	id, err:= strconv.ParseInt(idString, 10, 64)
	if err != nil{
		fmt.Println("Error parsing item ID", err)
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	item, err:=services.GetItem(id)
	if err!=nil{

		fmt.Println("Error getting item", err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, item)

}