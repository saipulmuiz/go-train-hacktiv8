package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Car struct {
	CarId string `json:"car_id"`
	Brand string `json:"brand"`
	Model string `json:"model"`
	Price int    `json:"price"`
}

var Cars []Car

func CreateCar(ctx *gin.Context) {
	var newCar Car

	err := ctx.ShouldBindJSON(&newCar)
	if err != nil {
		writeJsonResponse(ctx, http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	newCar.CarId = fmt.Sprintf("car-%d", len(Cars)+1)
	Cars = append(Cars, newCar)

	writeJsonResponse(ctx, http.StatusCreated, gin.H{
		"status": "success",
		"car":    newCar,
	})
}

func GetCar(ctx *gin.Context) {
	if len(Cars) == 0 {
		writeJsonResponse(ctx, http.StatusNotFound, gin.H{
			"error": "no data",
		})
		return
	}
	writeJsonResponse(ctx, http.StatusOK, gin.H{
		"status": "success",
		"data":   Cars,
	})
}

func DeleteCar(ctx *gin.Context) {
	param := ctx.Param("carId")
	isExist := false
	carIndex := 0

	for i, car := range Cars {
		if car.CarId == param {
			isExist = true
			carIndex = i
			break
		}
	}

	if !isExist {
		writeJsonResponse(ctx, http.StatusNotFound, gin.H{
			"error": fmt.Sprintf("data with id %s not found", param),
		})
		return
	}

	var newCars []Car
	newCars = append(newCars, Cars[:carIndex]...)
	newCars = append(newCars, Cars[carIndex+1:]...)

	Cars = newCars

	writeJsonResponse(ctx, http.StatusOK, newCars)
}

func writeJsonResponse(ctx *gin.Context, status int, payload interface{}) {
	ctx.JSON(status, payload)
}
