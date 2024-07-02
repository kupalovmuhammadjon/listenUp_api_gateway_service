package api

import (
	"api_gateway/api/handler"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	api := r.Group("api")

	h := handler.NewHandler()

	auth := api.Group("auth")
	auth.POST("/register")
	auth.POST("/login")
	auth.POST("/logout")
	auth.GET("/refresh-token")

	users := api.Group("users")
	users.GET("/{id}")
	users.PUT("/{id}")
	users.DELETE("/{id}")
	users.GET("/{id}/profile")
	users.PUT("/{id}/profile")

	compositions := api.Group("compositions")
	compositions.POST("/")
	compositions.GET("/{id}")
	compositions.PUT("/{id}")
	compositions.DELETE("/{id}")
	users.GET("/{id}/compositions")
	compositions.POST("/{id}/tracks")
	compositions.GET("/{id}/tracks")
	compositions.PUT("/{id}/tracks/{id}")
	compositions.DELETE("/{id}/tracks/{id}")
	compositions.POST("/{id}/publish")

	collaborations := api.Group("collaborations")
	collaborations.POST("/invite")
	collaborations.PUT("/invite/{id}/respond")
	compositions.GET("/{id}/collaborators")
	compositions.PUT("/{id}/collaborators/{id}")
	compositions.DELETE("/{id}/collaborators/{id}")
	compositions.POST("/{id}/comments")
	compositions.GET("/{id}/comments")

	discover := api.Group("discover")
	discover.GET("recommended")
	discover.GET("genres/{genre}")
	api.GET("search")
	compositions.POST("/{id}/like")
	compositions.DELETE("/{id}/like")
	compositions.POST("/{id}/listen")

	return r
}
