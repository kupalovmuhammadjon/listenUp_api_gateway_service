package handler

import (
	pb "api_gateway/genproto/podcasts"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *Handler) CreatePodcast(ctx *gin.Context) {
	req := pb.PodcastCreate{}

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
	resp, err := h.ClientPodcasts.CreatePodcast(nestedctx, &req)
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

func (h *Handler) GetPodcastById(ctx *gin.Context) {
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
	resp, err := h.ClientPodcasts.GetPodcastById(nestedctx, &req)
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

func (h *Handler) UpdatePodcast(ctx *gin.Context) {
	id := ctx.Param("id")
	if _, err := uuid.Parse(id); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "StatusBadRequest",
			"message": fmt.Sprintf("Error with getting Id from URL: %s", err),
		})
		log.Printf("Error with getting Id from URL: %s", err)
		return
	}
	req := pb.PodcastUpdate{}
	err := json.NewDecoder(ctx.Request.Body).Decode(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "StatusBadRequest",
			"message": fmt.Sprintf("Error with getting data from URL body: %s", err),
		})
		log.Printf("Error with getting data from URL body: %s", err)
		return
	}
	req.Id = id

	nestedctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()
	resp, err := h.ClientPodcasts.UpdatePodcast(nestedctx, &req)
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

func (h *Handler) DeletePodcast(ctx *gin.Context) {
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
	resp, err := h.ClientPodcasts.DeletePodcast(nestedctx, &req)
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

func (h *Handler) GetUserPodcasts(ctx *gin.Context) {
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
	resp, err := h.ClientPodcasts.GetUserPodcasts(nestedctx, &req)
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