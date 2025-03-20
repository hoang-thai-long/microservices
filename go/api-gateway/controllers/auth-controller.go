package controllers

import (
	"api-gateway/services"

	"github.com/gin-gonic/gin"
)

type AuthController interface {
	SignIn(ctx *gin.Context)
	SignUp(ctx *gin.Context)
	SignOut(ctx *gin.Context)
}

type controller struct {
	service services.MongoService
}
