package api

import (
	"api_gateway/api/handler"
	"api_gateway/config"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	api := r.Group("api")

	congif := config.Load()
	h := handler.NewHandler(congif)

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

	podcasts := api.Group("podcasts")
	podcasts.POST("/", h.CreatePodcast)
	podcasts.GET("/{id}", h.GetPodcastById)
	podcasts.PUT("/{id}", h.UpdatePodcast)
	podcasts.DELETE("/{id}", h.DeletePodcast)
	users.GET("/{id}/podcasts", h.GetUserPodcasts)
	podcasts.POST("/{id}/episodes", h.CreatePodcastEpisode)
	podcasts.GET("/{id}/episodes", h.GetEpisodesByPodcastId)
	podcasts.PUT("/{id}/episodes/{episodeid}", h.UpdateEpisode)
	podcasts.DELETE("/{id}/episodes/{episodeid}", h.DeleteEpisode)
	podcasts.POST("/{id}/publish", h.PublishPodcast)

	collaborations := api.Group("collaborations")
	collaborations.POST("/invite")
	collaborations.PUT("/invite/{id}/respond")
	podcasts.GET("/{id}/collaborators")
	podcasts.PUT("/{id}/collaborators/{userid}")
	podcasts.DELETE("/{id}/collaborators/{userid}")
	podcasts.POST("/{id}/comments")
	podcasts.GET("/{id}/comments")

	discover := api.Group("discover")
	discover.GET("recommended")
	discover.GET("genres/{genre}")
	api.GET("search")
	podcasts.POST("/{id}/like")
	podcasts.DELETE("/{id}/like")
	podcasts.POST("/{id}/listen")

	return r
}
