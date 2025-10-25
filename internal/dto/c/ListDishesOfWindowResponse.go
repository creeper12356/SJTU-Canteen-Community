package dto

import "SJTU-Canteen-Community/internal/model"

type ListDishesOfWindowResponse struct {
	Dishes []model.Dish `json:"dishes"`
	Total  int64        `json:"total"`
}
