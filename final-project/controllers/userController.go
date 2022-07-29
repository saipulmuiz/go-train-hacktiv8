package controllers

import (
	"final-project/helper"
	"final-project/models"
	"final-project/params"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
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
		badRequestResponse(ctx, err.Error())
		return
	}

	_, errCreate := govalidator.ValidateStruct(&user)
	if errCreate != nil {
		badRequestResponse(ctx, errCreate.Error())
		return
	}

	err = u.db.Create(&user).Error
	if err != nil {
		internalServerJsonResponse(ctx, err.Error())
		return
	}

	response := params.UserRegisterResponse{
		Id:       user.Id,
		Username: user.Username,
		Email:    user.Email,
		Age:      user.Age,
	}

	writeJsonResponse(ctx, http.StatusCreated, response)
}

func (u *UserController) Login(ctx *gin.Context) {
	var user models.User

	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		badRequestResponse(ctx, err.Error())
		return
	}

	var userDb models.User
	err = u.db.First(&userDb, "email=?", user.Email).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			noDataJsonResponse(ctx, err.Error())
			return
		}
		internalServerJsonResponse(ctx, err.Error())
		return
	}

	isValid := helper.ComparePassword(userDb.Password, user.Password)

	if !isValid {
		unauthorizeJsonResponse(ctx, "email or password is not match")
		return
	}

	token, err := helper.GenerateToken(userDb.Id, userDb.Email)
	if err != nil {
		internalServerJsonResponse(ctx, err.Error())
		return
	}

	writeJsonResponse(ctx, http.StatusOK, gin.H{
		"token": token,
	})
}

func (u *UserController) UpdateUser(ctx *gin.Context) {
	id, _ := ctx.Get("id")
	userId := uint(id.(float64))
	var user models.User

	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		badRequestResponse(ctx, err.Error())
		return
	}

	var userDb models.User
	err = u.db.First(&userDb, userId).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			noDataJsonResponse(ctx, err.Error())
			return
		}
		internalServerJsonResponse(ctx, err.Error())
		return
	}

	err = u.db.Model(&userDb).Updates(&user).Error
	if err != nil {
		badRequestResponse(ctx, err.Error())
		return
	}

	response := params.UserUpdateResponse{
		Id:        userDb.Id,
		Email:     userDb.Email,
		Username:  userDb.Username,
		Age:       userDb.Age,
		UpdatedAt: userDb.UpdatedAt,
	}

	writeJsonResponse(ctx, http.StatusOK, response)
}

func (u *UserController) DeleteUser(ctx *gin.Context) {
	id, _ := ctx.Get("id")
	userId := uint(id.(float64))

	var userDb models.User
	err := u.db.First(&userDb, userId).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			noDataJsonResponse(ctx, err.Error())
			return
		}
		internalServerJsonResponse(ctx, err.Error())
		return
	}

	err = u.db.Delete(&userDb).Error
	if err != nil {
		badRequestResponse(ctx, err.Error())
		return
	}

	writeJsonResponse(ctx, http.StatusOK, gin.H{
		"message": "Your account has been successfully deleted",
	})
}
