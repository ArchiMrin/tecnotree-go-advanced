package controllers

import (
	"github.com/ArchiMrin/tecnotree-go-advanced/ekart/services"
	"github.com/gin-gonic/gin"
)

func HandleLogin(c *gin.Context) {
	services.Login()
	c.JSON(200, gin.H{ //H in JSON
		"data": "Welcome to HandleLogin --- GO",
	})
}

func HandleResgister(c *gin.Context) {
	services.Register()
	c.JSON(200, gin.H{ //H in JSON
		"data": "Welcome to HandleRegister --- GO",
	})
}
func HandleGetProfile(c *gin.Context) {
	services.GetProfile()
	c.JSON(200, gin.H{ //H in JSON
		"data": "Welcome to HandleGetProfile --- GO",
	})
}
func HandleLogout(c *gin.Context) {
	services.Logout()
	c.JSON(200, gin.H{ //H in JSON
		"data": "Welcome to HandleLogout --- GO",
	})
}
