package controller

import (
	dto "SJTU-Canteen-Community/internal/dto/c"
	"SJTU-Canteen-Community/internal/service"
	"fmt"

	"github.com/gin-gonic/gin"
)

type WindowCommentController struct {
	CommentService     *service.CommentService
	ContentLikeService *service.ContentLikeService
}

func NewWindowCommentController(commentService *service.CommentService, contentLikeService *service.ContentLikeService) *WindowCommentController {
	return &WindowCommentController{
		CommentService:     commentService,
		ContentLikeService: contentLikeService,
	}
}

func (wcc *WindowCommentController) AddWindowComment(c *gin.Context) {
	var req dto.AddWindowCommentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	windowIDStr := c.Param("window_id")
	var windowID uint
	_, err := fmt.Sscanf(windowIDStr, "%d", &windowID)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid window_id"})
		return
	}
	req.WindowID = windowID

	windowCommentID, err := wcc.CommentService.AddWindowComment(&req, c.GetUint("user_id"))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"window_comment_id": windowCommentID})
}

func (wcc *WindowCommentController) LikeWindowComment(c *gin.Context) {
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
	err = wcc.ContentLikeService.LikeWindowComment(req.ContentID, userID, req.Weight)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Success"})
}
