package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func badRequestResponse(ctx *gin.Context, payload interface{}) {
	writeJsonResponse(ctx, http.StatusBadRequest, gin.H{
		"success": false,
		"error":   payload,
	})
}

func notFoundResponse(ctx *gin.Context, payload interface{}) {
	writeJsonResponse(ctx, http.StatusNotFound, gin.H{
		"success": false,
		"error":   payload,
	})
}

func noDataJsonResponse(ctx *gin.Context, err interface{}) {
	writeJsonResponse(ctx, http.StatusNotFound, gin.H{
		"success": false,
		"error":   err,
	})
}

func internalServerJsonResponse(ctx *gin.Context, err interface{}) {
	writeJsonResponse(ctx, http.StatusInternalServerError, gin.H{
		"success": false,
		"error":   err,
	})
}

func writeJsonResponse(ctx *gin.Context, status int, payload interface{}) {
	ctx.JSON(status, payload)
}

func unauthorizeJsonResponse(ctx *gin.Context, err interface{}) {
	writeJsonResponse(ctx, http.StatusNotFound, gin.H{
		"success": false,
		"error":   err,
	})
}
