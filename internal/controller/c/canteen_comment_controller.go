package controller

import (
	dto "SJTU-Canteen-Community/internal/dto/c"
	"SJTU-Canteen-Community/internal/service"
	"fmt"

	"github.com/gin-gonic/gin"
)

type CanteenCommentController struct {
	canteenCommentService *service.CommentService
	contentLikeService    *service.ContentLikeService
}

func NewCanteenCommentController(canteenCommentService *service.CommentService, contentLikeService *service.ContentLikeService) *CanteenCommentController {
	return &CanteenCommentController{
		canteenCommentService: canteenCommentService,
		contentLikeService:    contentLikeService,
	}
}

func (cc *CanteenCommentController) AddCanteenComment(c *gin.Context) {
	var req dto.AddCanteenCommentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	userID := c.GetUint("user_id")

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

func (cc *CanteenCommentController) LikeCanteenComment(c *gin.Context) {
	var req dto.LikeContentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	commentIDStr := c.Param("comment_id")
	_, err := fmt.Sscanf(commentIDStr, "%d", &req.ContentID)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid comment_id"})
		return
	}

	userID := c.GetUint("user_id")
	err = cc.contentLikeService.LikeCanteenComment(req.ContentID, userID, req.Weight)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Success"})
}
