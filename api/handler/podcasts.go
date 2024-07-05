package handler

import (
	pb "api_gateway/genproto/podcasts"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
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
			"message": fmt.Sprintf("Error with getting data from URL body: %s", err.Error()),
		})
		log.Printf("Error with getting data from URL body: %s", err.Error())
		return
	}
	nestedctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()
	resp, err := h.ClientPodcasts.CreatePodcast(nestedctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "StatusInternalServerError",
			"message": fmt.Sprintf("Error with request to podcasts service: %s", err.Error()),
		})
		log.Printf("Error with request to podcasts service: %s", err.Error())
		return
	}
	ctx.JSON(http.StatusAccepted, resp)
}

func (h *Handler) GetPodcastById(ctx *gin.Context) {
	id := ctx.Param("id")
	if _, err := uuid.Parse(id); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "StatusBadRequest",
			"message": fmt.Sprintf("Error with getting Id from URL: %s", err.Error()),
		})
		log.Printf("Error with getting Id from URL: %s", err.Error())
		return
	}
	req := pb.ID{Id: id}

	nestedctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()
	resp, err := h.ClientPodcasts.GetPodcastById(nestedctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "StatusInternalServerError",
			"message": fmt.Sprintf("Error with request to podcasts service: %s", err.Error()),
		})
		log.Printf("Error with request to podcasts service: %s", err.Error())
		return
	}
	ctx.JSON(http.StatusAccepted, resp)
}

func (h *Handler) UpdatePodcast(ctx *gin.Context) {
	id := ctx.Param("id")
	if _, err := uuid.Parse(id); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "StatusBadRequest",
			"message": fmt.Sprintf("Error with getting Id from URL: %s", err.Error()),
		})
		log.Printf("Error with getting Id from URL: %s", err.Error())
		return
	}
	req := pb.PodcastUpdate{}
	err := json.NewDecoder(ctx.Request.Body).Decode(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "StatusBadRequest",
			"message": fmt.Sprintf("Error with getting data from URL body: %s", err.Error()),
		})
		log.Printf("Error with getting data from URL body: %s", err.Error())
		return
	}
	req.Id = id

	nestedctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()
	resp, err := h.ClientPodcasts.UpdatePodcast(nestedctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "StatusInternalServerError",
			"message": fmt.Sprintf("Error with request to podcasts service: %s", err.Error()),
		})
		log.Printf("Error with request to podcasts service: %s", err.Error())
		return
	}
	ctx.JSON(http.StatusAccepted, resp)
}

func (h *Handler) DeletePodcast(ctx *gin.Context) {
	id := ctx.Param("id")
	if _, err := uuid.Parse(id); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "StatusBadRequest",
			"message": fmt.Sprintf("Error with getting Id from URL: %s", err.Error()),
		})
		log.Printf("Error with getting Id from URL: %s", err.Error())
		return
	}
	req := pb.ID{Id: id}

	nestedctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()
	resp, err := h.ClientPodcasts.DeletePodcast(nestedctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "StatusInternalServerError",
			"message": fmt.Sprintf("Error with request to podcasts service: %s", err.Error()),
		})
		log.Printf("Error with request to podcasts service: %s", err.Error())
		return
	}
	ctx.JSON(http.StatusAccepted, resp)
}

func (h *Handler) GetUserPodcasts(ctx *gin.Context) {
	id := ctx.Param("id")
	if _, err := uuid.Parse(id); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "StatusBadRequest",
			"message": fmt.Sprintf("Error with getting Id from URL: %s", err.Error()),
		})
		log.Printf("Error with getting Id from URL: %s", err.Error())
		return
	}

	limit := ctx.Param("limit")
	offset := ctx.Param("offset")
	if limit == "" || offset == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "StatusBadRequest",
			"message": fmt.Errorf("error with getting limit and offset from URL"),
		})
		log.Printf("Error with getting limit and offset from URL")
		return
	}

	limitInt, err := strconv.Atoi(limit)
	if limit == "" || offset == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "StatusBadRequest",
			"message": fmt.Errorf("error: limit not in int type: %s", limit),
		})
		log.Printf("Error: limit not in int type: %s", limit)
		return
	}

	offsetInt, err := strconv.Atoi(offset)
	if limit == "" || offset == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "StatusBadRequest",
			"message": fmt.Errorf("Error: limit not in int type: %s", offset),
		})
		log.Printf("Error: limit not in int type: %s", offset)
		return
	}

	req := pb.Filter{
		Id:     id,
		Limit:  int32(limitInt),
		Offset: int32(offsetInt),
	}

	nestedctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()
	resp, err := h.ClientPodcasts.GetUserPodcasts(nestedctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "StatusInternalServerError",
			"message": fmt.Sprintf("Error with request to podcasts service: %s", err.Error()),
		})
		log.Printf("Error with request to podcasts service: %s", err.Error())
		return
	}
	ctx.JSON(http.StatusAccepted, resp)
}

func (h *Handler) PublishPodcast(ctx *gin.Context) {
	id := ctx.Param("id")
	if _, err := uuid.Parse(id); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "StatusBadRequest",
			"message": fmt.Sprintf("Error with getting Id from URL: %s", err.Error()),
		})
		log.Printf("Error with getting Id from URL: %s", err.Error())
		return
	}

	req := pb.ID{
		Id: id,
	}

	nestedctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()
	resp, err := h.ClientPodcasts.PublishPodcast(nestedctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "StatusInternalServerError",
			"message": fmt.Sprintf("Error with request to podcasts service: %s", err.Error()),
		})
		log.Printf("Error with request to podcasts service: %s", err.Error())
		return
	}
	ctx.JSON(http.StatusAccepted, resp)
}
