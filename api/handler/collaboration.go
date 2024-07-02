package handler

import (
	pb "api_gateway/genproto/collaborations"
	"api_gateway/genproto/comments"
	pbc "api_gateway/genproto/comments"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
)

func (h *Handler) SendInvitation(ctx *gin.Context) {
	invitation := pb.CreateInvite{}
	err := json.NewDecoder(ctx.Request.Body).Decode(&invitation)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"Error": err,
			"Message": "Error while decoding",
		})
		log.Println("Error while decoding")
		return
	}

	tctx, cancel := context.WithTimeout(ctx, 5 * time.Second)
	defer cancel()

	h.ClientCollaboration.CreateInvitation(tctx, &invitation)
}

func (h *Handler) RepondInvitation(ctx *gin.Context) {
	collaboration := pb.CreateCollaboration{}
	err := json.NewDecoder(ctx.Request.Body).Decode(&collaboration)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"Error": err,
			"Message": "Error while decoding",
		})
		log.Println("Error while decoding ", err)
		return
	}

	id := ctx.Param("id")
	
	_, err = uuid.Parse(id)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"Error": err,
			"Message": "no id or invalid uuid",
		})
		log.Println("no id or invalid uuid ", err)
		return
	}
	collaboration.InvitationId = id

	tctx, cancel := context.WithTimeout(ctx, 5 * time.Second)
	defer cancel()

	collabId, err := h.ClientCollaboration.RespondInvitation(tctx, &collaboration)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"Error": err,
			"Message": "Error while responding and creating collaboration ",
		})
		log.Println("Error while responding and creating collaboration ", err)
		return
	}

	ctx.IndentedJSON(http.StatusCreated, gin.H{
		"Id": collabId,
	})
}

func (h *Handler) GetCollaboratorsByPodcastId(ctx *gin.Context){

	podcastId := ctx.Param("id")
	
	_, err := uuid.Parse(podcastId)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"Error": err,
			"Message": "no id or invalid uuid",
		})
		log.Println("no id or invalid uuid ", err)
		return
	}

	tctx, cancel := context.WithTimeout(ctx, 5 * time.Second)
	defer cancel()

	collaborators, err := h.ClientCollaboration.GetCollaboratorsByPodcastId(tctx, &pb.ID{Id: podcastId})
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"Error": err,
			"Message": "Error while getting collaborators by podcast_id",
		})
		log.Println("Error while getting collaborators by podcast_id ", err)
		return
	}

	ctx.JSON(http.StatusOK, collaborators)
}

func (h *Handler) UpdateCollaboratorByPodcastId(ctx *gin.Context){
	
	req := &pb.UpdateCollaborator{}
	err := json.NewDecoder(ctx.Request.Body).Decode(&req)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"Error": err,
			"Message": "Error while decoding",
		})
		log.Println("Error while decoding ", err)
		return
	}

	podcastId := ctx.Param("id")
	_, err = uuid.Parse(podcastId)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"Error": err,
			"Message": "no id or invalid uuid",
		})
		log.Println("no id or invalid uuid ", err)
		return
	}

	userId := ctx.Param("userId")
	_, err = uuid.Parse(userId)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"Error": err,
			"Message": "no userId or invalid uuid",
		})
		log.Println("no userId or invalid uuid ", err)
		return
	}

	req.PodcastId = podcastId
	req.UserId = userId

	tctx, cancel := context.WithTimeout(ctx, 5 * time.Second)
	defer cancel()

	_, err = h.ClientCollaboration.UpdateCollaboratorByPodcastId(tctx, req)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"Error": err,
			"Message": "error while updating collaborator by podcastId",
		})
		log.Println("error while updating collaborator by podcastId ", err)
		return
	}

}

func (h *Handler) DeleteCollaboratorByPodcastId(ctx *gin.Context){
	
	req := &pb.Ids{}
	podcastId := ctx.Param("id")
	_, err := uuid.Parse(podcastId)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"Error": err,
			"Message": "no id or invalid uuid",
		})
		log.Println("no id or invalid uuid ", err)
		return
	}

	userId := ctx.Param("userId")
	_, err = uuid.Parse(userId)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"Error": err,
			"Message": "no userId or invalid uuid",
		})
		log.Println("no userId or invalid uuid ", err)
		return
	}

	req.PodcastId = podcastId
	req.UserId = userId

	tctx, cancel := context.WithTimeout(ctx, 5 * time.Second)
	defer cancel()

	_, err = h.ClientCollaboration.DeleteCollaboratorByPodcastId(tctx, req)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"Error": err,
			"Message": "error while deleting collaborator by podcastId",
		})
		log.Println("error while deleting collaborator by podcastId ", err)
		return
	}

}

func (h *Handler) CreateCommentByPodcastId(ctx *gin.Context){
	
	req := &pbc.CreateComment{}
	err := json.NewDecoder(ctx.Request.Body).Decode(&req)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"Error": err,
			"Message": "Error while decoding",
		})
		log.Println("Error while decoding ", err)
		return
	}

	podcastId := ctx.Param("id")
	_, err = uuid.Parse(podcastId)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"Error": err,
			"Message": "no id or invalid uuid",
		})
		log.Println("no id or invalid uuid ", err)
		return
	}

	req.PodcastId = podcastId

	tctx, cancel := context.WithTimeout(ctx, 5 * time.Second)
	defer cancel()

	_, err = h.ClientComments.CreateCommentByPodcastId(tctx, req)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"Error": err,
			"Message": "error while posting comment by podcastId",
		})
		log.Println("error while posting comment by podcastId ", err)
		return
	}

}

func (h *Handler) GetCommentsByPodcastId(ctx *gin.Context){
	
	req := &pbc.ID{}
	err := json.NewDecoder(ctx.Request.Body).Decode(&req)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"Error": err,
			"Message": "Error while decoding",
		})
		log.Println("Error while decoding ", err)
		return
	}

	podcastId := ctx.Param("id")
	_, err = uuid.Parse(podcastId)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"Error": err,
			"Message": "no id or invalid uuid",
		})
		log.Println("no id or invalid uuid ", err)
		return
	}

	req.Id = podcastId

	tctx, cancel := context.WithTimeout(ctx, 5 * time.Second)
	defer cancel()

	comments, err := h.ClientComments.GetCommentsByPodcastId(tctx, req)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"Error": err,
			"Message": "error while posting comment by podcastId",
		})
		log.Println("error while posting comment by podcastId ", err)
		return
	}

	ctx.IndentedJSON(http.StatusOK, comments)

}