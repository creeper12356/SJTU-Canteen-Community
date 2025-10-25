package controller

import (
	dto "SJTU-Canteen-Community/internal/dto/c"
	"SJTU-Canteen-Community/internal/service"
	"fmt"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type CanteenCommentController struct {
	canteenCommentService *service.CommentService
}

func NewCanteenCommentController(canteenCommentService *service.CommentService) *CanteenCommentController {
	return &CanteenCommentController{
		canteenCommentService: canteenCommentService,
	}
}

func (cc *CanteenCommentController) AddCanteenComment(c *gin.Context) {
	var req dto.AddCanteenCommentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	session := sessions.Default(c)
	userID := session.Get("user_id").(uint)

	canteenIDStr := c.Param("canteen_id")
	_, err := fmt.Sscanf(canteenIDStr, "%d", &req.CanteenID)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid canteen_id"})
		return
	}

	commentID, err := cc.canteenCommentService.AddCanteenComment(&req, userID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"comment_id": commentID})
}
