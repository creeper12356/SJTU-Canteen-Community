package service

import (
	dto "SJTU-Canteen-Community/internal/dto/b"
	"SJTU-Canteen-Community/internal/model"
	"SJTU-Canteen-Community/internal/repository"
)

type DishService struct {
	repo *repository.DishRepository
}

func NewDishService(repo *repository.DishRepository) *DishService {
	return &DishService{repo: repo}
}

func (ds *DishService) MAddDishes(dto *dto.MAddDishesRequest) ([]model.Dish, error) {
	return ds.repo.MAddDishes(dto)
}
