package controllers

import (
	"assignment2/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type OrderController struct {
	db *gorm.DB
}

func NewOrderController(db *gorm.DB) *OrderController {
	return &OrderController{
		db: db,
	}
}

func (p *OrderController) GetAllOrder(ctx *gin.Context) {
	limit := ctx.Query("limit")
	limitInt := 10

	if limit != "" {
		l, err := strconv.Atoi(limit)
		if err == nil {
			limitInt = l
		}
	}

	var order []models.Order
	var total int

	err := p.db.Limit(limitInt).Find(&order).Count(&total).Preload("Items").Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			notFoundResponse(ctx, err.Error())
			return
		}
		badRequestResponse(ctx, err.Error())
		return
	}

	writeJsonResponse(ctx, http.StatusOK, gin.H{
		"success": true,
		"payload": order,
		"query": map[string]interface{}{
			"limit": limitInt,
			"total": total,
		},
	})
}

func (p *OrderController) CreateOrder(ctx *gin.Context) {
	var order models.Order

	err := ctx.ShouldBindJSON(&order)
	if err != nil {
		badRequestResponse(ctx, err.Error())
		return
	}

	p.db.Create(&order)
	for _, item := range order.Items {
		// item := &items[index]{orderId: order.OrderId}
		p.db.Create(item)
	}

	// for _, row := range order.Items {
	// 	sqlStr += "(?, ?, ?),"
	// 	vals = append(vals, row["v1"], row["v2"], row["v3"])
	// }
	// err = p.db.Create(&order).Error
	err = p.db.Model(&order).Association("Items").Append([]models.Item{}).Error
	if err != nil {
		badRequestResponse(ctx, err.Error())
		return
	}

	writeJsonResponse(ctx, http.StatusCreated, gin.H{
		"success": true,
		"message": "order created success",
	})
}

// func GetAllOrder(ctx *gin.Context) {
// 	limit := ctx.Query("limit")
// 	limitInt := 10

// 	if limit != "" {
// 		l, err := strconv.Atoi(limit)
// 		if err == nil {
// 			limitInt = l
// 		}
// 	}

// 	var people []models.Order
// 	var total int

// 	err := p.db.Limit(limitInt).Find(&people).Count(&total).Error
// 	if err != nil {
// 		if err.Error() == gorm.ErrRecordNotFound.Error() {
// 			notFoundResponse(ctx, err.Error())
// 			return
// 		}
// 		badRequestResponse(ctx, err.Error())
// 		return
// 	}

// 	writeJsonResponse(ctx, http.StatusOK, gin.H{
// 		"success": true,
// 		"payload": people,
// 		"query": map[string]interface{}{
// 			"limit": limitInt,
// 			"total": total,
// 		},
// 	})
// }

// func GetCar(ctx *gin.Context) {
// 	carID := ctx.Param("carID")
// 	condition := false
// 	var carData Car

// 	for i, car := range CarDatas {
// 		if carID == car.CarID {
// 			condition = true
// 			carData = CarDatas[i]
// 			break
// 		}
// 	}

// 	if !condition {
// 		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
// 			"error_status":  "Data Not Found",
// 		})
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, gin.H{
// 		"car": carData,
// 	})
// }

// func CreateCar(ctx *gin.Context) {
// 	var newCar Car

// 	if err := ctx.ShouldBindJSON(&newCar); err != nil {
// 		ctx.AbortWithError(http.StatusBadRequest, err)
// 		return
// 	}
// 	CarDatas = append(CarDatas, newCar)

// 	ctx.JSON(http.StatusCreated, gin.H{
// 		"car": newCar,
// 	})
// }

// func UpdateCar(ctx *gin.Context) {
// 	carID := ctx.Param("carID")
// 	condition := false
// 	var updatedCar Car

// 	if err := ctx.ShouldBindJSON(&updatedCar); err != nil {
// 		ctx.AbortWithError(http.StatusBadRequest, err)
// 		return
// 	}

// 	for i, car := range CarDatas {
// 		if carID == car.CarID {
// 			condition = true
// 			CarDatas[i] = updatedCar
// 			CarDatas[i].CarID = carID
// 			break
// 		}
// 	}

// 	if !condition {
// 		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
// 			"error_status":  "Data Not Found",
// 		})
// 		return
// 	}
// }

// func DeleteCar(ctx *gin.Context) {
// 	carID := ctx.Param("carID")
// 	condition := false
// 	var carIndex int

// 	for i, car := range CarDatas {
// 		if carID == car.CarID {
// 			condition = true
// 			carIndex = i
// 			break
// 		}
// 	}

// 	if !condition {
// 		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
// 			"error_status":  "Data Not Found",
// 		})
// 		return
// 	}

// 	copy(CarDatas[carIndex:], CarDatas[carIndex+1:])
// 	CarDatas[len(CarDatas)-1] = Car{}
// }
