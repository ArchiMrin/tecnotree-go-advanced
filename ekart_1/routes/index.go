package routes

import (
	"github.com/ArchiMrin/tecnotree-go-advanced/ekart_1/controllers"
	"github.com/gin-gonic/gin"
)

func Approutes(r *gin.Engine) {
	// user routes
	user := r.Group("/api/user")
	user.POST("/register", controllers.Register)
	user.POST("/login", controllers.Login)
	user.GET("/logout", controllers.Logout)
	user.GET("/profile/:id", controllers.Profile)

}

//Product Routes

func ProductRoutes(r *gin.Engine, p controllers.ProductController) {
	product := r.Group("/api/product")
	product.POST("/insert", p.InsertProduct)

}
