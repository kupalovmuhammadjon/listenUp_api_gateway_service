package api

import (
	"api_gateway/api/handler"
	"api_gateway/config"

	"github.com/gin-gonic/gin"
)

func NewRouter(cfg *config.Config) *gin.Engine {
	r := gin.Default()
	api := r.Group("api")

	h := handler.NewHandler(cfg)

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
	collaborations.POST("/invite", h.SendInvitation)
	collaborations.PUT("/invite/{id}/respond", h.RepondInvitation)
	compositions.GET("/{id}/collaborators", h.GetCollaboratorsByPodcastId)
	compositions.PUT("/{id}/collaborators/{userId}", h.UpdateCollaboratorByPodcastId)
	compositions.DELETE("/{id}/collaborators/{id}", h.DeleteCollaboratorByPodcastId)
	compositions.POST("/{id}/comments", h.CreateCommentByPodcastId)
	compositions.GET("/{id}/comments", h.GetCommentsByPodcastId)

	discover := api.Group("discover")
	discover.GET("recommended")
	discover.GET("genres/{genre}")
	api.GET("search")
	compositions.POST("/{id}/like")
	compositions.DELETE("/{id}/like")
	compositions.POST("/{id}/listen")

	return r
}
