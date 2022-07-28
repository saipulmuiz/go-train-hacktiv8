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

func (p *CommentController) CreateComment(ctx *gin.Context) {
	var comment models.Comment

	err := ctx.ShouldBindJSON(&comment)
	if err != nil {
		badRequestResponse(ctx, err.Error())
		return
	}

	id, _ := ctx.Get("id")
	comment.UserId = uint(id.(float64))

	err = p.db.Create(&comment).Error
	if err != nil {
		internalServerJsonResponse(ctx, err.Error())
		return
	}

	writeJsonResponse(ctx, http.StatusCreated, comment)
}

func (p *CommentController) GetComments(ctx *gin.Context) {
	var comments []models.Comment

	// id, _ := ctx.Get("id")
	err := p.db.Find(&comments).Error
	if err != nil {
		internalServerJsonResponse(ctx, err.Error())
		return
	}

	writeJsonResponse(ctx, http.StatusCreated, comments)
}

func (p *CommentController) UpdateComment(ctx *gin.Context) {
	commentId := ctx.Param("commentId")
	var comment models.Comment

	err := ctx.ShouldBindJSON(&comment)
	if err != nil {
		badRequestResponse(ctx, err.Error())
		return
	}

	var commentDb models.Comment
	err = p.db.First(&commentDb, "comment_id=?", commentId).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			noDataJsonResponse(ctx, err.Error())
			return
		}
		internalServerJsonResponse(ctx, err.Error())
		return
	}

	err = p.db.Model(&comment).Where("comment_id = ?", commentId).Updates(&comment).Error
	if err != nil {
		badRequestResponse(ctx, err.Error())
		return
	}

	writeJsonResponse(ctx, http.StatusOK, commentDb)
}

func (p *CommentController) DeleteComment(ctx *gin.Context) {
	commentId := ctx.Param("commentId")

	var commentDb models.Comment
	err := p.db.First(&commentDb, commentId).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			noDataJsonResponse(ctx, err.Error())
			return
		}
		internalServerJsonResponse(ctx, err.Error())
		return
	}

	err = p.db.Model(&commentDb).Delete(&commentDb, commentId).Error
	if err != nil {
		badRequestResponse(ctx, err.Error())
		return
	}

	writeJsonResponse(ctx, http.StatusCreated, gin.H{
		"message": "Your comment has been successfully deleted",
	})
}
