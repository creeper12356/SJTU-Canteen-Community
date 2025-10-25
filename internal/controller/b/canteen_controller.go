package controller

import (
	dto "SJTU-Canteen-Community/internal/dto/b"
	"SJTU-Canteen-Community/internal/service"
	"fmt"

	"github.com/gin-gonic/gin"
)

type CanteenController struct {
	service *service.CanteenService
}

func NewCanteenController(service *service.CanteenService) *CanteenController {
	return &CanteenController{service: service}
}

func (cc *CanteenController) AddCanteen(c *gin.Context) {
	var req dto.AddCanteenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	canteenID, err := cc.service.AddCanteen(&req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"canteen_id": canteenID})
}

func (cc *CanteenController) MAddWindowsToCanteen(c *gin.Context) {
	canteenIDStr := c.Param("canteen_id")
	var canteenID uint
	_, err := fmt.Sscanf(canteenIDStr, "%d", &canteenID)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid canteen_id"})
		return
	}

	var req dto.MAddWindowsToCanteenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	req.CanteenID = canteenID

	windows, err := cc.service.MAddWindowsToCanteen(&req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"windows": windows})
}

func (cc *CanteenController) MAddDishesToWindow(c *gin.Context) {
	windowIDStr := c.Param("window_id")
	var windowID uint
	_, err := fmt.Sscanf(windowIDStr, "%d", &windowID)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid window_id"})
		return
	}
	var req dto.MAddDishesToWindowRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	req.WindowID = windowID
	err = cc.service.MAddDishesToWindow(&req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Dishes added to window successfully"})
}
