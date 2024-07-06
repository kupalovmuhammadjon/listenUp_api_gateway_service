package middleware

import (
	"api_gateway/api/token"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		auth := ctx.GetHeader("Authorization")

		if auth == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Authorization header required",
			})
			return
		}

		valid, err := token.ValidateToken(auth)
		if err != nil || !valid {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "invalid token",
			})
			return
		}

		claims, err := token.ExtractClaims(auth)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "invalid token claims",
			})
			return
		}
		ctx.Set("claims", claims)
		ctx.Next()
	}
}
