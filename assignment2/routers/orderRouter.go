package routers

import (
	"assignment2/controllers"
	"assignment2/db"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()
	db := db.ConnectDB()

	orderController := controllers.NewOrderController(db)

	router.GET("/orders", orderController.GetAllOrder)
	router.POST("/orders", orderController.CreateOrder)

	// router.GET("/orders/:orderId", controllers.GetAllOrder)
	// router.PUT("/orders/:orderId", controllers.UpdateCar)
	// router.DELETE("/orders/:orderId", controllers.DeleteCar)

	return router
}
