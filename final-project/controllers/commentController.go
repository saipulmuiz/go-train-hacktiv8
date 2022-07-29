package controllers

import (
	"final-project/models"
	"final-project/params"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type CommentController struct {
	db *gorm.DB
}

func NewCommentController(db *gorm.DB) *CommentController {
	return &CommentController{db: db}
}

func (c *CommentController) PostComment(ctx *gin.Context) {
	var comment models.Comment

	err := ctx.ShouldBindJSON(&comment)
	if err != nil {
		badRequestResponse(ctx, err.Error())
		return
	}

	id, _ := ctx.Get("id")
	comment.UserId = uint(id.(float64))

	var photoDb models.Photo
	err = c.db.Model(&photoDb).First(&photoDb, comment.PhotoId).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			noDataJsonResponse(ctx, "Photo not found or has been deleted..")
			return
		}
		internalServerJsonResponse(ctx, err.Error())
		return
	}

	err = c.db.Create(&comment).Error
	if err != nil {
		internalServerJsonResponse(ctx, err.Error())
		return
	}

	response := params.CommentPostResponse{
		Id:        comment.Id,
		Message:   comment.Message,
		PhotoId:   comment.PhotoId,
		UserId:    comment.UserId,
		CreatedAt: comment.CreatedAt,
	}

	writeJsonResponse(ctx, http.StatusCreated, response)
}

func (c *CommentController) GetComments(ctx *gin.Context) {
	var comments []models.Comment

	err := c.db.Preload("User").Preload("Photo").Find(&comments).Error
	if err != nil {
		internalServerJsonResponse(ctx, err.Error())
		return
	}

	var resComments []params.CommentGetResponse
	for _, comment := range comments {
		resComment := params.CommentGetResponse{
			Id:        comment.Id,
			Message:   comment.Message,
			PhotoId:   comment.PhotoId,
			UserId:    comment.UserId,
			UpdatedAt: comment.UpdatedAt,
			CreatedAt: comment.CreatedAt,
			User: models.UserComment{
				Id:       comment.User.Id,
				Email:    comment.User.Email,
				Username: comment.User.Username,
			},
			Photo: models.PhotoComment{
				Id:       comment.Photo.Id,
				Title:    comment.Photo.Title,
				Caption:  comment.Photo.Caption,
				PhotoUrl: comment.Photo.PhotoUrl,
				UserId:   comment.Photo.UserId,
			},
		}
		resComments = append(resComments, resComment)
	}

	writeJsonResponse(ctx, http.StatusOK, resComments)
}

func (c *CommentController) UpdateComment(ctx *gin.Context) {
	commentId := ctx.Param("commentId")
	var comment models.Comment

	err := ctx.ShouldBindJSON(&comment)
	if err != nil {
		badRequestResponse(ctx, err.Error())
		return
	}
	id, _ := ctx.Get("id")
	userId := uint(id.(float64))

	var commentDb models.Comment
	err = c.db.First(&commentDb, commentId).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			noDataJsonResponse(ctx, err.Error())
			return
		}
		internalServerJsonResponse(ctx, err.Error())
		return
	}

	if commentDb.UserId != userId {
		unauthorizeJsonResponse(ctx, "You have not access to this data..")
		return
	}

	err = c.db.Model(&commentDb).Updates(&comment).Error
	if err != nil {
		badRequestResponse(ctx, err.Error())
		return
	}

	response := params.CommentUpdateResponse{
		Id:        commentDb.Id,
		Message:   commentDb.Message,
		PhotoId:   commentDb.PhotoId,
		UserId:    commentDb.UserId,
		UpdatedAt: commentDb.UpdatedAt,
	}

	writeJsonResponse(ctx, http.StatusOK, response)
}

func (c *CommentController) DeleteComment(ctx *gin.Context) {
	commentId := ctx.Param("commentId")
	id, _ := ctx.Get("id")
	userId := uint(id.(float64))

	var commentDb models.Comment
	err := c.db.First(&commentDb, commentId).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			noDataJsonResponse(ctx, err.Error())
			return
		}
		internalServerJsonResponse(ctx, err.Error())
		return
	}

	if commentDb.UserId != userId {
		unauthorizeJsonResponse(ctx, "You have not access to this data..")
		return
	}

	err = c.db.Delete(&commentDb).Error
	if err != nil {
		badRequestResponse(ctx, err.Error())
		return
	}

	writeJsonResponse(ctx, http.StatusOK, gin.H{
		"message": "Your comment has been successfully deleted",
	})
}
