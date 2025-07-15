package router

import (
	"user-api/controller"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine{
	r := gin.Default()

	r.POST("/register",controller.Register)
	r.POST("/login",controller.Login)
	return  r
}