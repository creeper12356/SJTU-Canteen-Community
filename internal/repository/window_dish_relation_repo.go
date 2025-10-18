package repository

import (
	dto "SJTU-Canteen-Community/internal/dto/b"
	"SJTU-Canteen-Community/internal/model"

	"gorm.io/gorm"
)

type WindowDishRelationRepository struct {
	db *gorm.DB
}

func NewWindowDishRelationRepository(db *gorm.DB) *WindowDishRelationRepository {
	return &WindowDishRelationRepository{db: db}
}

func (r *WindowDishRelationRepository) MAddDishesToWindow(dto *dto.MAddDishesToWindowRequest) error {
	var relations []model.WindowDishRelation
	for _, dish := range dto.Dishes {
		relation := model.WindowDishRelation{
			WindowID: dto.WindowID,
			DishID:   dish.ID,
			Price:    dish.Price,
		}
		relations = append(relations, relation)
	}

	if err := r.db.Create(&relations).Error; err != nil {
		return err
	}
	return nil
}
