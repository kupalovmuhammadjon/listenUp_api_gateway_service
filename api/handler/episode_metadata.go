package handler

import (
	pb "api_gateway/genproto/episode_metadata"
	"context"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func (h *Handler) GetTrendingPodcasts(c *gin.Context) {
	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"error": errors.Wrap(err, "invalid pagination parameters").Error()})
		log.Println(err)
		return
	}
	offset, err := strconv.Atoi(c.Query("offset"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"error": errors.Wrap(err, "invalid pagination parameters").Error()})
		log.Println(err)
		return
	}

	ctx, cancel := context.WithTimeout(c, time.Second*5)
	defer cancel()

	podcasts, err := h.ClientEpisodeMetadata.GetTrendingPodcasts(ctx, &pb.Pagination{
		Limit:  int64(limit),
		Offset: int64(offset),
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{"error": errors.Wrap(err, "failed to get trending podcasts").Error()})
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"Trending Podcasts": podcasts})
}

func (h *Handler) GetRecommendedPodcasts(c *gin.Context) {
	id := c.Param("userid")
	_, err := uuid.Parse(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"error": errors.Wrap(err, "invalid user id").Error()})
		log.Println(err)
		return
	}
	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"error": errors.Wrap(err, "invalid pagination parameters").Error()})
		log.Println(err)
		return
	}
	offset, err := strconv.Atoi(c.Query("offset"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": errors.Wrap(err, "invalid pagination parameters").Error()})
		log.Println(err)
		return
	}

	ctx, cancel := context.WithTimeout(c, time.Second*5)
	defer cancel()

	podcasts, err := h.ClientEpisodeMetadata.GetRecommendedPodcasts(ctx, &pb.IdPage{
		Id:         id,
		Pagination: &pb.Pagination{Limit: int64(limit), Offset: int64(offset)},
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{"error": errors.Wrap(err, "failed to get recommended podcasts").Error()})
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"Recommended Podcasts": podcasts})
}

func (h *Handler) GetPodcastsByGenre(c *gin.Context) {
	genres := c.QueryArray("genres")
	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"error": errors.Wrap(err, "invalid pagination parameters").Error()})
		log.Println(err)
		return
	}
	offset, err := strconv.Atoi(c.Query("offset"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"error": errors.Wrap(err, "invalid pagination parameters").Error()})
		log.Println(err)
		return
	}

	ctx, cancel := context.WithTimeout(c, time.Second*5)
	defer cancel()

	podcasts, err := h.ClientEpisodeMetadata.GetPodcastsByGenre(ctx, &pb.Filter{
		Genres:     genres,
		Pagination: &pb.Pagination{Limit: int64(limit), Offset: int64(offset)},
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{"error": errors.Wrap(err, "failed to get podcasts by genre").Error()})
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"Podcasts": podcasts})
}

func (h *Handler) SearchPodcast(c *gin.Context) {
	var title pb.Title
	err := c.ShouldBind(&title)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"error": errors.Wrap(err, "invalid data").Error()})
		log.Println(err)
		return
	}

	ctx, cancel := context.WithTimeout(c, time.Second*5)
	defer cancel()

	episode, err := h.ClientEpisodeMetadata.SearchEpisode(ctx, &title)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{"error": errors.Wrap(err, "failed to find episode").Error()})
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"Episode": episode})
}
