package repository

import (
	dto "SJTU-Canteen-Community/internal/dto/c"
	"SJTU-Canteen-Community/internal/model"

	"gorm.io/gorm"
)

type WindowCommentRepository struct {
	db *gorm.DB
}

func NewWindowCommentRepository(db *gorm.DB) *WindowCommentRepository {
	return &WindowCommentRepository{db: db}
}

func (r *WindowCommentRepository) AddWindowComment(dto *dto.AddWindowCommentRequest, userID uint) (uint, error) {
	windowComment := model.WindowComment{
		WindowID:   dto.WindowID,
		Content:    dto.Content,
		Rate:       dto.Rate,
		DishID:     dto.DishID,
		TotalPrice: dto.TotalPrice,
		UserID:     userID,
	}
	result := r.db.Create(&windowComment)
	return windowComment.ID, result.Error
}

func (r *WindowCommentRepository) CheckWindowCommentExists(commentID uint, forUpdate bool) (bool, error) {
	var count int64
	db := r.db
	if forUpdate {
		db = db.Set("gorm:query_option", "FOR UPDATE")
	}
	result := db.Model(&model.WindowComment{}).Where("id = ?", commentID).Limit(1).Count(&count)
	if result.Error != nil {
		return false, result.Error
	}
	return count > 0, nil
}
