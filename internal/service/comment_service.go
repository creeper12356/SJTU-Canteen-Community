package service

import (
	dto "SJTU-Canteen-Community/internal/dto/c"
	"SJTU-Canteen-Community/internal/repository"
)

type CommentService struct {
	canteenCommentRepo *repository.CanteenCommentRepository
}

func NewCommentService(canteenRepo *repository.CanteenCommentRepository) *CommentService {
	return &CommentService{
		canteenCommentRepo: canteenRepo,
	}
}

func (cs *CommentService) AddCanteenComment(dto *dto.AddCanteenCommentRequest, userID uint) (uint, error) {
	return cs.canteenCommentRepo.AddCanteenComment(dto, userID)
}
