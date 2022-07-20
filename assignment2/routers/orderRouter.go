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
	router.PUT("/orders/:orderId", orderController.UpdateOrder)
	router.DELETE("/orders/:orderId", orderController.DeleteOrder)

	return router
}
