package repository

import (
	dto "SJTU-Canteen-Community/internal/dto/c"
	"SJTU-Canteen-Community/internal/model"

	"gorm.io/gorm"
)

type CanteenCommentRepository struct {
	db *gorm.DB
}

func NewCanteenCommentRepository(db *gorm.DB) *CanteenCommentRepository {
	return &CanteenCommentRepository{db: db}
}

func (r *CanteenCommentRepository) AddCanteenComment(dto *dto.AddCanteenCommentRequest, userID uint) (uint, error) {
	canteenComment := model.CanteenComment{
		CanteenID: dto.CanteenID,
		Content:   dto.Content,
		Rate:      dto.Rate,
		UserID:    userID,
	}

	result := r.db.Create(&canteenComment)
	if result.Error != nil {
		return 0, result.Error
	}

	return canteenComment.ID, nil
}
