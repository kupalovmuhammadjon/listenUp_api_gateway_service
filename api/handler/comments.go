package handler

import (
	pbc "api_gateway/genproto/comments"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
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

	req := &pbc.CommentFilter{}

	podcastId := ctx.Param("id")
	_, err := uuid.Parse(podcastId)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"Error":   err,
			"Message": "no id or invalid uuid",
		})
		log.Println("no id or invalid uuid ", err)
		return
	}
	req.Id = podcastId

	l := ctx.Param("limit")
	if len(l) == 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Error":   err,
			"Message": "no limit in query",
		})
		log.Println("no limit in query ", err)
		return
	}

	limit, err := strconv.Atoi(l)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Error":   err,
			"Message": "invalid limit in query",
		})
		log.Println("invalid limit in query ", err)
		return
	}

	o := ctx.Param("offset")
	if len(l) == 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Error":   err,
			"Message": "no offset in query",
		})
		log.Println("no offset in query ", err)
		return
	}
	offset, err := strconv.Atoi(o)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Error":   err,
			"Message": "invalid offset in query",
		})
		log.Println("invalid offset in query ", err)
		return
	}

	req.Limit = int32(limit)
	req.Offset = int32(offset)

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
