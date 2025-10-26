package repository

import (
	dto_b "SJTU-Canteen-Community/internal/dto/b"
	dto_c "SJTU-Canteen-Community/internal/dto/c"
	"SJTU-Canteen-Community/internal/model"

	"gorm.io/gorm"
)

type CanteenRepository struct {
	db *gorm.DB
}

func NewCanteenRepository(db *gorm.DB) *CanteenRepository {
	return &CanteenRepository{db: db}
}

func (r *CanteenRepository) AddCanteen(dto *dto_b.AddCanteenRequest) (uint, error) {
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
	result := r.db.Model(&model.Canteen{}).Where("id = ?", canteenID).Limit(1).Count(&count)
	if result.Error != nil {
		return false, result.Error
	}
	return count > 0, nil
}

func (r *CanteenRepository) ListCanteens(dto *dto_c.ListCanteensRequest) (dto_c.ListCanteensResponse, error) {
	var canteens []model.Canteen
	var total int64

	result := r.db.Model(&model.Canteen{}).Count(&total)
	if result.Error != nil {
		return dto_c.ListCanteensResponse{}, result.Error
	}

	result = r.db.Offset(dto.Page * dto.PageSize).Limit(dto.PageSize).Find(&canteens)
	if result.Error != nil {
		return dto_c.ListCanteensResponse{}, result.Error
	}
	return dto_c.ListCanteensResponse{
		Canteens: canteens,
		Total:    total,
	}, nil
}
