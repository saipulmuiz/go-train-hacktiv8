package main

import (
	"restapi/controllers"
	"restapi/db"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	db := db.ConnectDB()

	personController := controllers.NewPersonController(db)

	router.POST("/person", personController.CreatePerson)
	router.GET("/person", personController.GetAllPeople)

	PORT := ":4444"
	router.Run(PORT)
}
