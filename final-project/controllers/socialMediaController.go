package controllers

import (
	"final-project/models"
	"final-project/params"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type SocialMediaController struct {
	db *gorm.DB
}

func NewSocialMediaController(db *gorm.DB) *SocialMediaController {
	return &SocialMediaController{db: db}
}

func (s *SocialMediaController) CreateSocialMedia(ctx *gin.Context) {
	var socialMediaReq params.SocialMediaRequest

	err := ctx.ShouldBindJSON(&socialMediaReq)
	if err != nil {
		badRequestResponse(ctx, err.Error())
		return
	}

	id, _ := ctx.Get("id")
	socialMediaReq.UserId = uint(id.(float64))

	_, errCreate := govalidator.ValidateStruct(&socialMediaReq)
	if errCreate != nil {
		badRequestResponse(ctx, errCreate.Error())
		return
	}

	socialMedia := models.SocialMedia{
		Name:           socialMediaReq.Name,
		SocialMediaUrl: socialMediaReq.SocialMediaUrl,
		UserId:         socialMediaReq.UserId,
	}

	err = s.db.Create(&socialMedia).Error
	if err != nil {
		internalServerJsonResponse(ctx, err.Error())
		return
	}

	response := params.SocialMediaPostResponse{
		Id:             socialMedia.Id,
		Name:           socialMedia.Name,
		SocialMediaUrl: socialMedia.SocialMediaUrl,
		UserId:         socialMedia.UserId,
		CreatedAt:      socialMedia.CreatedAt,
	}

	writeJsonResponse(ctx, http.StatusCreated, response)
}

func (s *SocialMediaController) GetSocialMedias(ctx *gin.Context) {
	var socialMedias []models.SocialMedia

	err := s.db.Preload("User").Find(&socialMedias).Error
	if err != nil {
		internalServerJsonResponse(ctx, err.Error())
		return
	}

	var resSocialMedias []params.SocialMediaGetResponse
	for _, social := range socialMedias {
		resSocmed := params.SocialMediaGetResponse{
			Id:             social.Id,
			Name:           social.Name,
			SocialMediaUrl: social.SocialMediaUrl,
			UserId:         social.UserId,
			CreatedAt:      social.CreatedAt,
			UpdatedAt:      social.UpdatedAt,
			User: models.UserSocialMedia{
				Id:              social.UserId,
				Username:        social.User.Username,
				ProfileImageUrl: "Profile Image URL",
			},
		}
		resSocialMedias = append(resSocialMedias, resSocmed)
	}

	writeJsonResponse(ctx, http.StatusOK, gin.H{
		"social_medias": resSocialMedias,
	})
}

func (s *SocialMediaController) UpdateSocialMedia(ctx *gin.Context) {
	socialMediaId := ctx.Param("socialMediaId")
	var socialMediaReq params.SocialMediaRequest

	err := ctx.ShouldBindJSON(&socialMediaReq)
	if err != nil {
		badRequestResponse(ctx, err.Error())
		return
	}
	id, _ := ctx.Get("id")
	userId := uint(id.(float64))

	var socialMediaDb models.SocialMedia
	err = s.db.First(&socialMediaDb, socialMediaId).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			noDataJsonResponse(ctx, err.Error())
			return
		}
		internalServerJsonResponse(ctx, err.Error())
		return
	}

	if socialMediaDb.UserId != userId {
		unauthorizeJsonResponse(ctx, "You have not access to this data..")
		return
	}

	_, errUpdate := govalidator.ValidateStruct(&socialMediaReq)
	if errUpdate != nil {
		badRequestResponse(ctx, errUpdate.Error())
		return
	}

	socialMedia := models.SocialMedia{
		Name:           socialMediaReq.Name,
		SocialMediaUrl: socialMediaReq.SocialMediaUrl,
		UserId:         socialMediaReq.UserId,
	}

	err = s.db.Model(&socialMediaDb).Updates(&socialMedia).Error
	if err != nil {
		badRequestResponse(ctx, err.Error())
		return
	}

	response := params.SocialMediaUpdateResponse{
		Id:             socialMediaDb.Id,
		Name:           socialMediaDb.Name,
		SocialMediaUrl: socialMediaDb.SocialMediaUrl,
		UserId:         socialMediaDb.UserId,
		UpdatedAt:      socialMediaDb.UpdatedAt,
	}

	writeJsonResponse(ctx, http.StatusOK, response)
}

func (s *SocialMediaController) DeleteSocialMedia(ctx *gin.Context) {
	socialMediaId := ctx.Param("socialMediaId")
	id, _ := ctx.Get("id")
	userId := uint(id.(float64))

	var socialMediaDb models.SocialMedia
	err := s.db.First(&socialMediaDb, socialMediaId).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			noDataJsonResponse(ctx, err.Error())
			return
		}
		internalServerJsonResponse(ctx, err.Error())
		return
	}

	if socialMediaDb.UserId != userId {
		unauthorizeJsonResponse(ctx, "You have not access to this data..")
		return
	}

	err = s.db.Delete(&socialMediaDb).Error
	if err != nil {
		badRequestResponse(ctx, err.Error())
		return
	}

	writeJsonResponse(ctx, http.StatusOK, gin.H{
		"message": "Your social media has been successfully deleted",
	})
}
