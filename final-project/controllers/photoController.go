package controllers

import (
	"final-project/models"
	"final-project/params"
	"net/http"

	valid "github.com/asaskevich/govalidator"
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
	var photoReq params.PhotoRequest

	err := ctx.ShouldBindJSON(&photoReq)
	if err != nil {
		badRequestResponse(ctx, err.Error())
		return
	}

	id, _ := ctx.Get("id")
	photoReq.UserId = uint(id.(float64))

	_, errCreate := valid.ValidateStruct(&photoReq)
	if errCreate != nil {
		badRequestResponse(ctx, errCreate.Error())
		return
	}

	photo := models.Photo{
		Title:    photoReq.Title,
		Caption:  photoReq.Caption,
		PhotoUrl: photoReq.PhotoUrl,
		UserId:   photoReq.UserId,
	}

	err = p.db.Create(&photo).Error
	if err != nil {
		internalServerJsonResponse(ctx, err.Error())
		return
	}

	response := params.PhotoPostResponse{
		Id:        photo.Id,
		Title:     photo.Title,
		Caption:   photo.Caption,
		PhotoUrl:  photo.PhotoUrl,
		UserId:    photo.UserId,
		CreatedAt: photo.CreatedAt,
	}

	writeJsonResponse(ctx, http.StatusCreated, response)
}

func (p *PhotoController) GetPhotos(ctx *gin.Context) {
	var photos []models.Photo

	err := p.db.Preload("User").Find(&photos).Error
	if err != nil {
		internalServerJsonResponse(ctx, err.Error())
		return
	}

	var resPhotos []params.PhotoGetResponse
	for _, photo := range photos {
		resPhoto := params.PhotoGetResponse{
			Id:        photo.Id,
			Title:     photo.Title,
			Caption:   photo.Caption,
			PhotoUrl:  photo.PhotoUrl,
			UserId:    photo.UserId,
			CreatedAt: photo.CreatedAt,
			UpdatedAt: photo.UpdatedAt,
			User: models.UserPhoto{
				Email:    photo.User.Email,
				Username: photo.User.Username,
			},
		}
		resPhotos = append(resPhotos, resPhoto)
	}

	writeJsonResponse(ctx, http.StatusOK, resPhotos)
}

func (p *PhotoController) UpdatePhoto(ctx *gin.Context) {
	photoId := ctx.Param("photoId")
	var photoReq params.PhotoRequest

	err := ctx.ShouldBindJSON(&photoReq)
	if err != nil {
		badRequestResponse(ctx, err.Error())
		return
	}
	id, _ := ctx.Get("id")
	userId := uint(id.(float64))

	var photoDb models.Photo
	err = p.db.First(&photoDb, photoId).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			noDataJsonResponse(ctx, err.Error())
			return
		}
		internalServerJsonResponse(ctx, err.Error())
		return
	}

	if photoDb.UserId != userId {
		unauthorizeJsonResponse(ctx, "You have not access to this data..")
		return
	}

	_, errUpdate := valid.ValidateStruct(&photoReq)
	if errUpdate != nil {
		badRequestResponse(ctx, errUpdate.Error())
		return
	}

	photo := models.Photo{
		Title:    photoReq.Title,
		Caption:  photoReq.Caption,
		PhotoUrl: photoReq.PhotoUrl,
		UserId:   photoReq.UserId,
	}

	err = p.db.Model(&photoDb).Updates(&photo).Error
	if err != nil {
		badRequestResponse(ctx, err.Error())
		return
	}

	response := params.PhotoUpdateResponse{
		Id:        photoDb.Id,
		Title:     photoDb.Title,
		Caption:   photoDb.Caption,
		PhotoUrl:  photoDb.PhotoUrl,
		UserId:    photoDb.UserId,
		UpdatedAt: photoDb.UpdatedAt,
	}

	writeJsonResponse(ctx, http.StatusOK, response)
}

func (p *PhotoController) DeletePhoto(ctx *gin.Context) {
	photoId := ctx.Param("photoId")
	id, _ := ctx.Get("id")
	userId := uint(id.(float64))

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

	if photoDb.UserId != userId {
		unauthorizeJsonResponse(ctx, "You have not access to this data..")
		return
	}

	err = p.db.Delete(&photoDb).Error
	if err != nil {
		badRequestResponse(ctx, err.Error())
		return
	}

	writeJsonResponse(ctx, http.StatusOK, gin.H{
		"message": "Your photo has been successfully deleted",
	})
}
