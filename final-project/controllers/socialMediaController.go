package controllers

import (
	"final-project/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type SocialMediaController struct {
	db *gorm.DB
}

func NewSocialMediaController(db *gorm.DB) *SocialMediaController {
	return &SocialMediaController{db: db}
}

func (p *SocialMediaController) CreateSocialMedia(ctx *gin.Context) {
	var socialMedia models.SocialMedia

	err := ctx.ShouldBindJSON(&socialMedia)
	if err != nil {
		badRequestResponse(ctx, err.Error())
		return
	}

	id, _ := ctx.Get("id")
	socialMedia.UserId = uint(id.(float64))

	err = p.db.Create(&socialMedia).Error
	if err != nil {
		internalServerJsonResponse(ctx, err.Error())
		return
	}

	writeJsonResponse(ctx, http.StatusCreated, socialMedia)
}

func (p *SocialMediaController) GetSocialMedias(ctx *gin.Context) {
	var socialMedias []models.SocialMedia

	// id, _ := ctx.Get("id")
	err := p.db.Find(&socialMedias).Error
	if err != nil {
		internalServerJsonResponse(ctx, err.Error())
		return
	}

	writeJsonResponse(ctx, http.StatusCreated, gin.H{
		"social_medias": socialMedias,
	})
}

func (p *SocialMediaController) UpdateSocialMedia(ctx *gin.Context) {
	socialMediaId := ctx.Param("socialMediaId")
	var socialMedia models.SocialMedia

	err := ctx.ShouldBindJSON(&socialMedia)
	if err != nil {
		badRequestResponse(ctx, err.Error())
		return
	}

	var socialMediaDb models.SocialMedia
	err = p.db.First(&socialMediaDb, "socialMedia_id=?", socialMediaId).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			noDataJsonResponse(ctx, err.Error())
			return
		}
		internalServerJsonResponse(ctx, err.Error())
		return
	}

	err = p.db.Model(&socialMedia).Where("socialMedia_id = ?", socialMediaId).Updates(&socialMedia).Error
	if err != nil {
		badRequestResponse(ctx, err.Error())
		return
	}

	writeJsonResponse(ctx, http.StatusOK, socialMediaDb)
}

func (p *SocialMediaController) DeleteSocialMedia(ctx *gin.Context) {
	socialMediaId := ctx.Param("socialMediaId")

	var socialMediaDb models.SocialMedia
	err := p.db.First(&socialMediaDb, socialMediaId).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			noDataJsonResponse(ctx, err.Error())
			return
		}
		internalServerJsonResponse(ctx, err.Error())
		return
	}

	err = p.db.Model(&socialMediaDb).Delete(&socialMediaDb, socialMediaId).Error
	if err != nil {
		badRequestResponse(ctx, err.Error())
		return
	}

	writeJsonResponse(ctx, http.StatusCreated, gin.H{
		"success": true,
		"message": "Your social media has been successfully deleted",
	})
}
