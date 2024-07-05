package handler

import (
	pb "api_gateway/genproto/user_interactions"
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func (h *Handler) LikeEpisodeOfPodcast(c *gin.Context) {
	var interaction pb.InteractEpisode
	err := c.ShouldBind(&interaction)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"error": errors.Wrap(err, "invalid data").Error()})
		log.Println(err)
		return
	}

	ctx, cancel := context.WithTimeout(c, time.Second*5)
	defer cancel()

	id, err := h.ClientUserInteractions.LikeEpisodeOfPodcast(ctx, &interaction)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{"error": errors.Wrap(err, "failed to like episode").Error()})
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"New Interaction ID": id})
}

func (h *Handler) DeleteLikeFromEpisodeOfPodcast(c *gin.Context) {
	var ids pb.DeleteLike
	err := c.ShouldBind(&ids)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"error": errors.Wrap(err, "invalid data").Error()})
		log.Println(err)
		return
	}

	ctx, cancel := context.WithTimeout(c, time.Second*5)
	defer cancel()

	success, err := h.ClientUserInteractions.DeleteLikeFromEpisodeOfPodcast(ctx, &ids)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{"error": errors.Wrap(err, "failed to dislike episode").Error()})
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"Successful": success.Success})
}

func (h *Handler) ListenEpisodeOfPodcast(c *gin.Context) {
	var interaction pb.InteractEpisode
	err := c.ShouldBind(&interaction)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"error": errors.Wrap(err, "invalid data").Error()})
		log.Println(err)
		return
	}

	ctx, cancel := context.WithTimeout(c, time.Second*5)
	defer cancel()

	id, err := h.ClientUserInteractions.ListenEpisodeOfPodcast(ctx, &interaction)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{"error": errors.Wrap(err, "failed to listen to episode").Error()})
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"New Interaction ID": id})
}
