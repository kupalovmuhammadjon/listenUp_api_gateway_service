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

	users := api.Group("users")
	users.GET("/{id}", h.GetUserByID)
	users.PUT("/{id}", h.UpdateUser)
	users.DELETE("/{id}", h.DeleteUser)
	users.GET("/{id}/profile", h.GetUserProfile)
	users.PUT("/{id}/profile", h.UpdateUserProfile)

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
	collaborations.POST("/invite", h.SendInvitation)
	collaborations.PUT("/invite/{id}/respond", h.RepondInvitation)
	podcasts.GET("/{id}/collaborators", h.GetCollaboratorsByPodcastId)
	podcasts.PUT("/{id}/collaborators/{userid}", h.UpdateCollaboratorByPodcastId)
	podcasts.DELETE("/{id}/collaborators/{userid}", h.DeleteCollaboratorByPodcastId)
	podcasts.POST("/{id}/comments", h.CreateCommentByPodcastId)
	podcasts.GET("/{id}/comments", h.GetCommentsByPodcastId)

	discover := api.Group("discover")
	discover.GET("trending", h.GetTrendingPodcasts)
	discover.GET("recommended", h.GetRecommendedPodcasts)
	discover.GET("genres/{genre}", h.GetPodcastsByGenre)
	api.GET("search", h.SearchPodcast)
	podcasts.POST("/{id}/like", h.LikeEpisodeOfPodcast)
	podcasts.DELETE("/{id}/like", h.DeleteLikeFromEpisodeOfPodcast)
	podcasts.POST("/{id}/listen", h.ListenEpisodeOfPodcast)

	return r
}
