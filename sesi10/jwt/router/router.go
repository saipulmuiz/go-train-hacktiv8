package router

import (
	"jwt/controller"
	"jwt/database"
	"jwt/middleware"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	router := gin.Default()

	db := database.ConnectDB()
	userController := controller.NewUserController(db)
	productController := controller.NewProductController(db)

	router.Use(middleware.Auth())
	userGroup := router.Group("/users")
	{
		userGroup.POST("/register", userController.Register)
		userGroup.POST("/login", userController.Login)
	}

	productGroup := router.Group("/products")
	{

		// productGroup.Use(middleware.Auth())
		productGroup.GET("/", productController.GetProducts)
		productGroup.POST("/", productController.CreateProduct)
	}

	return router
}
