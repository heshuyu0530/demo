package controller

import (
	"net/http"
	"user-api/model"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
	    c.JSON(http.StatusBadRequest, gin.H{"error": "无效请求"})
		return
	}
	if !model.Register(user) {
		c.JSON(http.StatusConflict, gin.H{"error": "用户已存在"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "注册成功"})
}

func Login(c *gin.Context){
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效请求"})
		return
	}
	if !model.Login(user) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "用户名或密码错误"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "登录成功"})
}