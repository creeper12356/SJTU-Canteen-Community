package repository

import (
	dto "SJTU-Canteen-Community/internal/dto/b"
	"SJTU-Canteen-Community/internal/model"

	"gorm.io/gorm"
)

type WindowRepository struct {
	db *gorm.DB
}

func NewWindowRepository(db *gorm.DB) *WindowRepository {
	return &WindowRepository{db: db}
}

func (r *WindowRepository) AddWindowsToCanteen(dto *dto.MAddWindowsToCanteenRequest) ([]model.Window, error) {
	var windows []model.Window
	for _, windowDTO := range dto.Windows {
		window := model.Window{
			Name:        windowDTO.Name,
			Description: windowDTO.Description,
			PictureURL:  windowDTO.PictureURL,
			CanteenID:   dto.CanteenID,
		}
		windows = append(windows, window)
	}
	if err := r.db.Create(&windows).Error; err != nil {
		return nil, err
	}
	return windows, nil
}

func (r *WindowRepository) CheckWindowExists(windowID uint) (bool, error) {
	var count int64
	result := r.db.Model(&model.Window{}).Where("id = ?", windowID).Count(&count)
	if result.Error != nil {
		return false, result.Error
	}
	return count > 0, nil
}
