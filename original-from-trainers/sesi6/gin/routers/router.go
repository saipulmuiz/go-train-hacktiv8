package routers

import (
	"gin/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.POST("/cars", controllers.CreateCar)
	router.GET("/cars", controllers.GetCar)
	router.DELETE("/cars/:carId", controllers.DeleteCar)
	return router
}
