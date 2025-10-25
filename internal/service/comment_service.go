package service

import (
	dto "SJTU-Canteen-Community/internal/dto/c"
	"SJTU-Canteen-Community/internal/repository"
	"fmt"
)

type CommentService struct {
	canteenCommentRepo *repository.CanteenCommentRepository
	canteenRepo        *repository.CanteenRepository
}

func NewCommentService(canteenCommentRepo *repository.CanteenCommentRepository, canteenRepo *repository.CanteenRepository) *CommentService {
	return &CommentService{
		canteenCommentRepo: canteenCommentRepo,
		canteenRepo:        canteenRepo,
	}
}

func (cs *CommentService) AddCanteenComment(dto *dto.AddCanteenCommentRequest, userID uint) (uint, error) {
	exists, err := cs.canteenRepo.CheckCanteenExists(dto.CanteenID)
	if err != nil {
		return 0, err
	}
	if !exists {
		return 0, fmt.Errorf("canteen with id %d does not exist", dto.CanteenID)
	}
	return cs.canteenCommentRepo.AddCanteenComment(dto, userID)
}
