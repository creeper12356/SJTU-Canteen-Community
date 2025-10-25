package repository

import (
	dto "SJTU-Canteen-Community/internal/dto/b"
	dto_c "SJTU-Canteen-Community/internal/dto/c"
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

func (r *WindowDishRelationRepository) ListDishIDsOfWindow(dto dto_c.ListDishesOfWindowRequest) ([]uint, int64, error) {
	var relations []model.WindowDishRelation
	result := r.db.Where("window_id = ?", dto.WindowID).Offset(dto.Page * dto.PageSize).Limit(dto.PageSize).Find(&relations)

	if result.Error != nil {
		return nil, 0, result.Error
	}

	var dishIDs []uint
	for _, relation := range relations {
		dishIDs = append(dishIDs, relation.DishID)
	}

	var total int64
	r.db.Model(&model.WindowDishRelation{}).Where("window_id = ?", dto.WindowID).Count(&total)

	return dishIDs, total, nil
}
