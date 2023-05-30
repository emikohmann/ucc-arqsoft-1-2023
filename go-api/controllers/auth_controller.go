package controllers

import (
	"github.com/gin-gonic/gin"
	"go-api/domain"
	"go-api/services"
	"net/http"
)

func Auth(ctx *gin.Context) {
	var user domain.User

	if err := ctx.BindJSON(&user); err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	authResponse, err := services.Auth(user)
	if err != nil {
		ctx.Status(http.StatusForbidden)
		return
	}

	ctx.JSON(http.StatusOK, authResponse)
}
