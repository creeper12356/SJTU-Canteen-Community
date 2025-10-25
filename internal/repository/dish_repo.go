package repository

import (
	dto "SJTU-Canteen-Community/internal/dto/b"
	"SJTU-Canteen-Community/internal/model"

	"gorm.io/gorm"
)

type DishRepository struct {
	db *gorm.DB
}

func NewDishRepository(db *gorm.DB) *DishRepository {
	return &DishRepository{db: db}
}

func (r *DishRepository) MAddDishes(dto *dto.MAddDishesRequest) ([]model.Dish, error) {
	var dishModels []model.Dish
	for _, dish := range dto.Dishes {
		dishModels = append(dishModels, model.Dish{
			Name:        dish.Name,
			PictureURL:  dish.PictureURL,
			Description: dish.Description,
		})
	}
	if err := r.db.Create(&dishModels).Error; err != nil {
		return nil, err
	}
	return dishModels, nil
}

func (r *DishRepository) MCheckAllDishesExist(dishIDs []uint) (bool, error) {
	var count int64
	result := r.db.Model(&model.Dish{}).Where("id IN ?", dishIDs).Count(&count)
	if result.Error != nil {
		return false, result.Error
	}
	return count == int64(len(dishIDs)), nil
}

func (r *DishRepository) ListDishesByIDs(dishIDs []uint) ([]model.Dish, error) {
	var dishes []model.Dish
	result := r.db.Where("id IN ?", dishIDs).Find(&dishes)
	if result.Error != nil {
		return nil, result.Error
	}
	return dishes, nil
}
