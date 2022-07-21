package controller

import (
	"jwt/helper"
	"jwt/models"
	"jwt/params"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController struct {
	db *gorm.DB
}

func NewUserController(db *gorm.DB) *UserController {
	return &UserController{db: db}
}

func (u *UserController) Register(ctx *gin.Context) {
	var user models.User

	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		badRequestJsonResponse(ctx, err.Error())
		return
	}

	_, errCreate := govalidator.ValidateStruct(&user)
	if errCreate != nil {
		badRequestJsonResponse(ctx, errCreate.Error())
		return
	}

	err = u.db.Create(&user).Error
	if err != nil {
		internalServerJsonResponse(ctx, err.Error())
		return
	}

	response := params.UserRegisterResponse{
		ID:       user.ID,
		FullName: user.FullName,
		Email:    user.Email,
	}

	writeJsonResponse(ctx, http.StatusCreated, gin.H{
		"success": true,
		"payload": response,
	})
}

func (u *UserController) Login(ctx *gin.Context) {
	var user models.User

	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		badRequestJsonResponse(ctx, err.Error())
		return
	}

	var userDB models.User

	err = u.db.First(&userDB, "email=?", user.Email).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			noDataJsonResponse(ctx, err.Error())
			return
		}
		internalServerJsonResponse(ctx, err.Error())
		return
	}

	isValid := helper.ComparePassword(userDB.Password, user.Password)

	if !isValid {
		unauthorizeJsonResponse(ctx, "username / password is not match")
		return
	}

	token, err := helper.GenerateToken(userDB.ID, userDB.Email)
	if err != nil {
		internalServerJsonResponse(ctx, err.Error())
		return
	}

	writeJsonResponse(ctx, http.StatusOK, gin.H{
		"success": true,
		"payload": map[string]string{
			"token": token,
		},
	})
}
