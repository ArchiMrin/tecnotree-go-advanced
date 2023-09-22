package controllers

import (
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	//services.GetRegister()
	c.JSON(200, gin.H{
		"data": "Welcome to gin--REGISTER",
	})

}

func Login(c *gin.Context) {
	//services.GetLogin()
	c.JSON(200, gin.H{
		"data": "Welcome to gin--LOGIN",
	})

}

func Logout(c *gin.Context) {
	//services.GetLogout()
	c.JSON(200, gin.H{
		"data": "Welcome to gin--LOGOUT",
	})

}

func Profile(c *gin.Context) {
	//services.GetProfile()
	c.JSON(200, gin.H{
		"data": "Welcome to gin--PROFILE",
	})

}
