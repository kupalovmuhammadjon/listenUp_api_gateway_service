package handler

import (
	pbc "api_gateway/genproto/comments"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *Handler) CreateCommentByPodcastId(ctx *gin.Context) {

	req := &pbc.CreateComment{}
	err := json.NewDecoder(ctx.Request.Body).Decode(&req)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"Error":   err,
			"Message": "Error while decoding",
		})
		log.Println("Error while decoding ", err)
		return
	}

	podcastId := ctx.Param("id")
	_, err = uuid.Parse(podcastId)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"Error":   err,
			"Message": "no id or invalid uuid",
		})
		log.Println("no id or invalid uuid ", err)
		return
	}

	req.PodcastId = podcastId

	tctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	_, err = h.ClientComments.CreateCommentByPodcastId(tctx, req)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"Error":   err,
			"Message": "error while posting comment by podcastId",
		})
		log.Println("error while posting comment by podcastId ", err)
		return
	}

}

func (h *Handler) GetCommentsByPodcastId(ctx *gin.Context) {

	req := &pbc.ID{}
	err := json.NewDecoder(ctx.Request.Body).Decode(&req)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"Error":   err,
			"Message": "Error while decoding",
		})
		log.Println("Error while decoding ", err)
		return
	}

	podcastId := ctx.Param("id")
	_, err = uuid.Parse(podcastId)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"Error":   err,
			"Message": "no id or invalid uuid",
		})
		log.Println("no id or invalid uuid ", err)
		return
	}

	req.Id = podcastId

	tctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	comments, err := h.ClientComments.GetCommentsByPodcastId(tctx, req)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"Error":   err,
			"Message": "error while posting comment by podcastId",
		})
		log.Println("error while posting comment by podcastId ", err)
		return
	}

	ctx.IndentedJSON(http.StatusOK, comments)

}
