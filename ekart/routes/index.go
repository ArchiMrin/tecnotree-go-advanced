package routes

import (
	"github.com/ArchiMrin/tecnotree-go-advanced/ekart/controllers"
	"github.com/gin-gonic/gin"
)

func AppRoutes(r *gin.Engine) { // passing the gin engine
	//user routes
	user := r.Group("/api/user")
	user.POST("/register", controllers.HandleResgister)

	user.POST("/login", controllers.HandleLogin)
	user.GET("/logout", controllers.HandleLogout)
	user.GET("/profile/:id", controllers.HandleGetProfile)
	//product routes
	//product := r.Group("/api/product")
	//product.Post("/add", controllers.HandleResgister)
	//product.Post("/edit", controllers.HandleLogin)
	//product.Post("/search", controllers.HandleLogout)
	//product.Post("/product/:id", controllers.HandleGetProfile)

}
