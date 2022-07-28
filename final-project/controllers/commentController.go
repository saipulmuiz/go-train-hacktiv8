package controllers

import (
	"final-project/models"
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

	err = c.db.Create(&comment).Error
	if err != nil {
		internalServerJsonResponse(ctx, err.Error())
		return
	}

	writeJsonResponse(ctx, http.StatusCreated, comment)
}

func (c *CommentController) GetComments(ctx *gin.Context) {
	var comments []models.Comment

	// id, _ := ctx.Get("id")
	err := c.db.Find(&comments).Error
	if err != nil {
		internalServerJsonResponse(ctx, err.Error())
		return
	}

	writeJsonResponse(ctx, http.StatusCreated, comments)
}

func (c *CommentController) UpdateComment(ctx *gin.Context) {
	commentId := ctx.Param("commentId")
	var comment models.Comment

	err := ctx.ShouldBindJSON(&comment)
	if err != nil {
		badRequestResponse(ctx, err.Error())
		return
	}

	var commentDb models.Comment
	err = c.db.First(&commentDb, "comment_id=?", commentId).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			noDataJsonResponse(ctx, err.Error())
			return
		}
		internalServerJsonResponse(ctx, err.Error())
		return
	}

	err = c.db.Model(&comment).Where("comment_id = ?", commentId).Updates(&comment).Error
	if err != nil {
		badRequestResponse(ctx, err.Error())
		return
	}

	writeJsonResponse(ctx, http.StatusOK, commentDb)
}

func (c *CommentController) DeleteComment(ctx *gin.Context) {
	commentId := ctx.Param("commentId")

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

	err = c.db.Model(&commentDb).Delete(&commentDb, commentId).Error
	if err != nil {
		badRequestResponse(ctx, err.Error())
		return
	}

	writeJsonResponse(ctx, http.StatusCreated, gin.H{
		"message": "Your comment has been successfully deleted",
	})
}
