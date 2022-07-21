package controller

import (
	"jwt/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ProductController struct {
	db *gorm.DB
}

func NewProductController(db *gorm.DB) *ProductController {
	return &ProductController{db: db}
}

func (p *ProductController) CreateProduct(ctx *gin.Context) {
	var product models.Product

	err := ctx.ShouldBindJSON(&product)
	if err != nil {
		badRequestJsonResponse(ctx, err.Error())
		return
	}

	id, _ := ctx.Get("id")
	product.UserID = uint(id.(float64))

	err = p.db.Create(&product).Error
	if err != nil {
		internalServerJsonResponse(ctx, err.Error())
		return
	}

	writeJsonResponse(ctx, http.StatusCreated, gin.H{
		"success": true,
		"payload": product,
	})
}

func (p *ProductController) GetProducts(ctx *gin.Context) {
	var products []models.Product

	err := p.db.Find(&products).Error
	if err != nil {
		internalServerJsonResponse(ctx, err.Error())
		return
	}

	writeJsonResponse(ctx, http.StatusCreated, gin.H{
		"success": true,
		"payload": products,
	})
}
