package handler

import (
	pbmetadat "api_gateway/genproto/episode_metadata"
	pb "api_gateway/genproto/episodes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *Handler) CreatePodcastEpisode(ctx *gin.Context) {
	id := ctx.Param("id")
	if _, err := uuid.Parse(id); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "StatusBadRequest",
			"message": fmt.Sprintf("Error with getting Id from URL: %s", err),
		})
		log.Printf("Error with getting Id from URL: %s", err)
		return
	}

	req := pb.EpisodeCreate{PodcastId: id}
	err := json.NewDecoder(ctx.Request.Body).Decode(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "StatusBadRequest",
			"message": fmt.Sprintf("Error with getting data from URL body: %s", err),
		})
		log.Printf("Error with getting data from URL body: %s", err)
		return
	}

	nestedctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()
	resp, err := h.ClientEpisodes.CreatePodcastEpisode(nestedctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "StatusInternalServerError",
			"message": fmt.Sprintf("Error with request to podcasts service: %s", err),
		})
		log.Printf("Error with request to podcasts service: %s", err)
		return
	}

	// create Epiosode metadata
	req2 := pbmetadat.EpisodeMetadata{
		EpisodeId: resp.Id,
		PodcastId: req.PodcastId,
		Genre:     req.Genre,
		Tags:      req.Tags,
	}
	nestedctx1, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()
	_, err = h.ClientEpisodeMetadata.CreateEpisodeMetaData(nestedctx1, &req2)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "StatusInternalServerError",
			"message": fmt.Sprintf("Error with creating episode metadate: %s", err),
		})
		log.Printf("Error with creating episode metadate: %s", err)
		return
	}
	ctx.JSON(http.StatusAccepted, resp)
}

func (h *Handler) GetEpisodesByPodcastId(ctx *gin.Context) {
	id := ctx.Param("id")
	if _, err := uuid.Parse(id); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "StatusBadRequest",
			"message": fmt.Sprintf("Error with getting Id from URL: %s", err),
		})
		log.Printf("Error with getting Id from URL: %s", err)
		return
	}

	req := pb.ID{Id: id}

	nestedctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()
	resp, err := h.ClientEpisodes.GetEpisodesByPodcastId(nestedctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "StatusInternalServerError",
			"message": fmt.Sprintf("Error with request to podcasts service: %s", err),
		})
		log.Printf("Error with request to podcasts service: %s", err)
		return
	}
	ctx.JSON(http.StatusAccepted, resp)
}

func (h *Handler) UpdateEpisode(ctx *gin.Context) {
	podcastId := ctx.Param("id")
	if _, err := uuid.Parse(podcastId); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "StatusBadRequest",
			"message": fmt.Sprintf("Error with getting Id from URL: %s", err),
		})
		log.Printf("Error with getting Id from URL: %s", err)
		return
	}
	episodeId := ctx.Param("episodeid")
	if _, err := uuid.Parse(episodeId); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "StatusBadRequest",
			"message": fmt.Sprintf("Error with getting Id from URL: %s", err),
		})
		log.Printf("Error with getting Id from URL: %s", err)
		return
	}

	req := pb.IDs{
		PodcastId: podcastId,
		EpisodeId: episodeId,
	}
	err := json.NewDecoder(ctx.Request.Body).Decode(&req.Episode)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "StatusBadRequest",
			"message": fmt.Sprintf("Error with getting data from URL body: %s", err),
		})
		log.Printf("Error with getting data from URL body: %s", err)
		return
	}

	nestedctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()
	resp, err := h.ClientEpisodes.UpdateEpisode(nestedctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "StatusInternalServerError",
			"message": fmt.Sprintf("Error with request to podcasts service: %s", err),
		})
		log.Printf("Error with request to podcasts service: %s", err)
		return
	}
	ctx.JSON(http.StatusAccepted, resp)
}

func (h *Handler) DeleteEpisode(ctx *gin.Context) {
	podcastId := ctx.Param("id")
	if _, err := uuid.Parse(podcastId); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "StatusBadRequest",
			"message": fmt.Sprintf("Error with getting Id from URL: %s", err),
		})
		log.Printf("Error with getting Id from URL: %s", err)
		return
	}
	episodeId := ctx.Param("episodeid")
	if _, err := uuid.Parse(episodeId); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "StatusBadRequest",
			"message": fmt.Sprintf("Error with getting Id from URL: %s", err),
		})
		log.Printf("Error with getting Id from URL: %s", err)
		return
	}

	req := pb.IDsForDelete{
		PodcastId: podcastId,
		EpisodeId: episodeId,
	}

	nestedctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()
	resp, err := h.ClientEpisodes.DeleteEpisode(nestedctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "StatusInternalServerError",
			"message": fmt.Sprintf("Error with request to podcasts service: %s", err),
		})
		log.Printf("Error with request to podcasts service: %s", err)
		return
	}
	ctx.JSON(http.StatusAccepted, resp)
}

func (h *Handler) PublishPodcast(ctx *gin.Context) {
	id := ctx.Param("id")
	if _, err := uuid.Parse(id); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "StatusBadRequest",
			"message": fmt.Sprintf("Error with getting Id from URL: %s", err),
		})
		log.Printf("Error with getting Id from URL: %s", err)
		return
	}

	req := pb.ID{
		Id: id,
	}

	nestedctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()
	resp, err := h.ClientEpisodes.PublishPodcast(nestedctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "StatusInternalServerError",
			"message": fmt.Sprintf("Error with request to podcasts service: %s", err),
		})
		log.Printf("Error with request to podcasts service: %s", err)
		return
	}
	ctx.JSON(http.StatusAccepted, resp)
}
