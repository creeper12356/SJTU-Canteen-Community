package controller

import (
	dto "SJTU-Canteen-Community/internal/dto/c"
	"SJTU-Canteen-Community/internal/service"
	"fmt"

	"github.com/gin-gonic/gin"
)

type CanteenController struct {
	canteenService *service.CanteenService
}

func NewCanteenController(canteenService *service.CanteenService) *CanteenController {
	return &CanteenController{canteenService: canteenService}
}

func (cc *CanteenController) ListCanteens(ctx *gin.Context) {
	var req dto.ListCanteensRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	resp, err := cc.canteenService.ListCanteens(&req)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, resp)
}

func (cc *CanteenController) ListWindowsOfCanteen(ctx *gin.Context) {
	canteenIDStr := ctx.Param("canteen_id")
	var canteenID uint
	_, err := fmt.Sscanf(canteenIDStr, "%d", &canteenID)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid canteen_id"})
		return
	}

	var req dto.ListWindowsOfCanteenRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	req.CanteenID = canteenID

	resp, err := cc.canteenService.ListWindowsOfCanteen(&req)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, resp)
}

func (cc *CanteenController) ListDishesOfWindow(ctx *gin.Context) {
	windowIDStr := ctx.Param("window_id")
	var windowID uint
	_, err := fmt.Sscanf(windowIDStr, "%d", &windowID)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid window_id"})
		return
	}

	var req dto.ListDishesOfWindowRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	req.WindowID = windowID

	resp, err := cc.canteenService.ListDishesOfWindow(&req)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, resp)
}
