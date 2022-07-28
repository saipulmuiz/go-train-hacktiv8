package controllers

import (
	"final-project/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type PhotoController struct {
	db *gorm.DB
}

func NewPhotoController(db *gorm.DB) *PhotoController {
	return &PhotoController{db: db}
}

func (p *PhotoController) CreatePhoto(ctx *gin.Context) {
	var photo models.Photo

	err := ctx.ShouldBindJSON(&photo)
	if err != nil {
		badRequestResponse(ctx, err.Error())
		return
	}

	id, _ := ctx.Get("id")
	photo.UserId = uint(id.(float64))

	err = p.db.Create(&photo).Error
	if err != nil {
		internalServerJsonResponse(ctx, err.Error())
		return
	}

	writeJsonResponse(ctx, http.StatusCreated, photo)
}

func (p *PhotoController) GetPhotos(ctx *gin.Context) {
	var photos []models.Photo

	// id, _ := ctx.Get("id")
	err := p.db.Find(&photos).Error
	if err != nil {
		internalServerJsonResponse(ctx, err.Error())
		return
	}

	writeJsonResponse(ctx, http.StatusCreated, photos)
}

func (p *PhotoController) UpdatePhoto(ctx *gin.Context) {
	photoId := ctx.Param("photoId")
	var photo models.Photo

	err := ctx.ShouldBindJSON(&photo)
	if err != nil {
		badRequestResponse(ctx, err.Error())
		return
	}

	var photoDb models.Photo
	err = p.db.First(&photoDb, "photo_id=?", photoId).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			noDataJsonResponse(ctx, err.Error())
			return
		}
		internalServerJsonResponse(ctx, err.Error())
		return
	}

	err = p.db.Model(&photo).Where("photo_id = ?", photoId).Updates(&photo).Error
	if err != nil {
		badRequestResponse(ctx, err.Error())
		return
	}

	writeJsonResponse(ctx, http.StatusOK, photoDb)
}

func (p *PhotoController) DeletePhoto(ctx *gin.Context) {
	photoId := ctx.Param("photoId")

	var photoDb models.Photo
	err := p.db.First(&photoDb, photoId).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			noDataJsonResponse(ctx, err.Error())
			return
		}
		internalServerJsonResponse(ctx, err.Error())
		return
	}

	err = p.db.Model(&photoDb).Delete(&photoDb, photoId).Error
	if err != nil {
		badRequestResponse(ctx, err.Error())
		return
	}

	writeJsonResponse(ctx, http.StatusCreated, gin.H{
		"message": "Your photo has been successfully deleted",
	})
}
