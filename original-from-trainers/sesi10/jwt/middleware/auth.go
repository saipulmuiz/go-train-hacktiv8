package middleware

import (
	"jwt/helper"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		headerToken := ctx.Request.Header.Get("Authorization")
		if headerToken == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"error":   "UNAUTHORIZED",
			})
			return
		}

		bearer := strings.HasPrefix(headerToken, "Bearer")
		if !bearer {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"error":   "UNAUTHORIZED",
			})
			return
		}

		bearerToken := strings.Split(headerToken, "Bearer ")[1]

		verify, err := helper.ValidateToken(bearerToken)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"error":   err.Error(),
			})
			return
		}
		data := verify.(jwt.MapClaims)

		ctx.Set("id", data["id"])
		ctx.Set("email", data["email"])
		ctx.Next()
	}
}
