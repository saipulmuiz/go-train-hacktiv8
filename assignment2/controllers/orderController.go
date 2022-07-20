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

func (oc *OrderController) GetAllOrder(ctx *gin.Context) {
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

	err := oc.db.Limit(limitInt).Preload("Items").Find(&order).Count(&total).Error
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

func (oc *OrderController) CreateOrder(ctx *gin.Context) {
	var order models.Order

	err := ctx.ShouldBindJSON(&order)
	if err != nil {
		badRequestResponse(ctx, err.Error())
		return
	}

	oc.db.Create(&order)
	for _, item := range order.Items {
		item.OrderId = int(order.OrderId)
		oc.db.Create(&item)
	}
	err = oc.db.Model(&order).Association("Items").Error
	if err != nil {
		badRequestResponse(ctx, err.Error())
		return
	}

	writeJsonResponse(ctx, http.StatusCreated, gin.H{
		"success": true,
		"message": "Order created successfully",
	})
}

func (oc *OrderController) UpdateOrder(ctx *gin.Context) {
	orderId := ctx.Param("orderId")
	var order models.Order

	err := ctx.ShouldBindJSON(&order)
	if err != nil {
		badRequestResponse(ctx, err.Error())
		return
	}

	var orderDb models.Order
	err = oc.db.First(&orderDb, "order_id=?", orderId).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			noDataJsonResponse(ctx, err.Error())
			return
		}
		internalServerJsonResponse(ctx, err.Error())
		return
	}

	err = oc.db.Model(&order).Where("order_id = ?", orderId).Updates(&order).Error
	if err != nil {
		badRequestResponse(ctx, err.Error())
		return
	}

	writeJsonResponse(ctx, http.StatusOK, gin.H{
		"success": true,
		"message": "Order updated successfully",
	})
}

func (oc *OrderController) DeleteOrder(ctx *gin.Context) {
	orderId := ctx.Param("orderId")

	var orderDb models.Order
	err := oc.db.First(&orderDb, orderId).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			noDataJsonResponse(ctx, err.Error())
			return
		}
		internalServerJsonResponse(ctx, err.Error())
		return
	}

	err = oc.db.Debug().Model(&orderDb).Delete(&orderDb, orderId).Delete(orderDb.Items, "order_id = ?", orderId).Error
	if err != nil {
		badRequestResponse(ctx, err.Error())
		return
	}

	writeJsonResponse(ctx, http.StatusCreated, gin.H{
		"success": true,
		"message": "Order deleted success",
	})
}
