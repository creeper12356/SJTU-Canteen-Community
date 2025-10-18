package repository

import (
	dto "SJTU-Canteen-Community/internal/dto/b"
	"SJTU-Canteen-Community/internal/model"

	"gorm.io/gorm"
)

type CanteenRepository struct {
	db *gorm.DB
}

func NewCanteenRepository(db *gorm.DB) *CanteenRepository {
	return &CanteenRepository{db: db}
}

func (r *CanteenRepository) AddCanteen(dto *dto.AddCanteenRequest) (uint, error) {
	canteen := model.Canteen{
		Name:        dto.Name,
		PictureURL:  dto.PictureURL,
		Description: dto.Description,
	}

	result := r.db.Create(&canteen)
	if result.Error != nil {
		return 0, result.Error
	}

	return canteen.ID, nil
}

func (r *CanteenRepository) CheckCanteenExists(canteenID uint) (bool, error) {
	var count int64
	result := r.db.Model(&model.Canteen{}).Where("id = ?", canteenID).Count(&count)
	if result.Error != nil {
		return false, result.Error
	}
	return count > 0, nil
}
