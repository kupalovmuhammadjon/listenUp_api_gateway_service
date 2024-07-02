package handler

import (
	pb "api_gateway/genproto/episode_metadata"
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func (h *Handler) GetTrendingPodcasts(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, time.Second*5)
	defer cancel()

	podcasts, err := h.ClientEpisodeMetadata.GetTrendingPodcasts(ctx, &pb.Void{})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": errors.Wrap(err, "failed to get trending podcasts")})
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"Trending Podcasts": podcasts})
}

func (h *Handler) GetRecommendedPodcasts(c *gin.Context) {
	id := c.Param("user_id")
	_, err := uuid.Parse(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": errors.Wrap(err, "invalid user id")})
		log.Println(err)
		return
	}

	ctx, cancel := context.WithTimeout(c, time.Second*5)
	defer cancel()

	podcasts, err := h.ClientEpisodeMetadata.GetRecommendedPodcasts(ctx, &pb.ID{Id: id})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": errors.Wrap(err, "failed to get recommended podcasts")})
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"Recommended Podcasts": podcasts})
}

func (h *Handler) GetPodcastsByGenre(c *gin.Context) {
	var genres pb.Genres
	err := c.ShouldBind(&genres)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": errors.Wrap(err, "invalid data")})
		log.Println(err)
		return
	}

	ctx, cancel := context.WithTimeout(c, time.Second*5)
	defer cancel()

	podcasts, err := h.ClientEpisodeMetadata.GetPodcastsByGenre(ctx, &genres)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": errors.Wrap(err, "failed to get podcasts by genre")})
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"Podcasts": podcasts})
}

func (h *Handler) SearchPodcast(c *gin.Context) {
	var titles pb.Title
	err := c.ShouldBind(&titles)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": errors.Wrap(err, "invalid data")})
		log.Println(err)
		return
	}

	ctx, cancel := context.WithTimeout(c, time.Second*5)
	defer cancel()

	podcasts, err := h.ClientEpisodeMetadata.SearchPodcast(ctx, &titles)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": errors.Wrap(err, "failed to search podcasts")})
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"Podcasts": podcasts})
}
