package controller

import (
	dto "SJTU-Canteen-Community/internal/dto/b"
	"SJTU-Canteen-Community/internal/service"

	"github.com/gin-gonic/gin"
)

type DishController struct {
	service *service.DishService
}

func NewDishController(service *service.DishService) *DishController {
	return &DishController{service: service}
}

func (dc *DishController) MAddDishes(c *gin.Context) {
	var req dto.MAddDishesRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	dishes, err := dc.service.MAddDishes(&req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, dishes)
}
