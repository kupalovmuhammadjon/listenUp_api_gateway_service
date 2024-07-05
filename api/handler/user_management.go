package handler

import (
	pb "api_gateway/genproto/user"
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func (h *Handler) GetUserByID(c *gin.Context) {
	id := c.Param("user_id")
	_, err := uuid.Parse(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"error": errors.Wrap(err, "invalid user id").Error()})
		log.Println(err)
		return
	}

	ctx, cancel := context.WithTimeout(c, time.Second*5)
	defer cancel()

	user, err := h.ClientUserManagement.GetUserByID(ctx, &pb.ID{Id: id})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{"error": errors.Wrap(err, "failed to get user").Error()})
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"User": user})
}

func (h *Handler) UpdateUser(c *gin.Context) {
	id := c.Param("user_id")
	_, err := uuid.Parse(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"error": errors.Wrap(err, "invalid user id").Error()})
		log.Println(err)
		return
	}

	var user pb.User
	err = c.ShouldBind(&user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"error": errors.Wrap(err, "invalid data").Error()})
		log.Println(err)
		return
	}

	ctx, cancel := context.WithTimeout(c, time.Second*5)
	defer cancel()

	_, err = h.ClientUserManagement.UpdateUser(ctx, &user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{"error": errors.Wrap(err, "failed to update user").Error()})
		log.Println(err)
		return
	}

	c.JSON(http.StatusNoContent, "User updated successfully")
}

func (h *Handler) DeleteUser(c *gin.Context) {
	id := c.Param("user_id")
	_, err := uuid.Parse(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"error": errors.Wrap(err, "invalid user id").Error()})
		log.Println(err)
		return
	}

	ctx, cancel := context.WithTimeout(c, time.Second*5)
	defer cancel()

	_, err = h.ClientUserManagement.DeleteUser(ctx, &pb.ID{Id: id})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{"error": errors.Wrap(err, "failed to delete user").Error()})
		log.Println(err)
		return
	}

	c.JSON(http.StatusNoContent, "User deleted successfully")
}

func (h *Handler) GetUserProfile(c *gin.Context) {
	id := c.Param("user_id")
	_, err := uuid.Parse(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"error": errors.Wrap(err, "invalid user id").Error()})
		log.Println(err)
		return
	}

	ctx, cancel := context.WithTimeout(c, time.Second*5)
	defer cancel()

	profile, err := h.ClientUserManagement.GetUserProfile(ctx, &pb.ID{Id: id})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{"error": errors.Wrap(err, "failed to get user profile").Error()})
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"User Profile": profile})
}

func (h *Handler) UpdateUserProfile(c *gin.Context) {
	id := c.Param("user_id")
	_, err := uuid.Parse(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"error": errors.Wrap(err, "invalid user id").Error()})
		log.Println(err)
		return
	}

	var profile pb.Profile
	err = c.ShouldBind(&profile)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"error": errors.Wrap(err, "invalid data").Error()})
		log.Println(err)
		return
	}

	ctx, cancel := context.WithTimeout(c, time.Second*5)
	defer cancel()

	_, err = h.ClientUserManagement.UpdateUserProfile(ctx, &profile)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{"error": errors.Wrap(err, "failed to update user profile").Error()})
		log.Println(err)
		return
	}

	c.JSON(http.StatusNoContent, "User profile deleted successfully")
}
