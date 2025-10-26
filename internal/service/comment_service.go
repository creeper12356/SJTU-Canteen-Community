package service

import (
	dto "SJTU-Canteen-Community/internal/dto/c"
	"SJTU-Canteen-Community/internal/repository"
	"fmt"
)

type CommentService struct {
	canteenCommentRepo     *repository.CanteenCommentRepository
	canteenRepo            *repository.CanteenRepository
	windowRepo             *repository.WindowRepository
	windowDishRelationRepo *repository.WindowDishRelationRepository
	windowCommentRepo      *repository.WindowCommentRepository
}

func NewCommentService(
	canteenCommentRepo *repository.CanteenCommentRepository,
	canteenRepo *repository.CanteenRepository,
	windowRepo *repository.WindowRepository,
	windowDishRelationRepo *repository.WindowDishRelationRepository,
	windowCommentRepo *repository.WindowCommentRepository) *CommentService {
	return &CommentService{
		canteenCommentRepo:     canteenCommentRepo,
		canteenRepo:            canteenRepo,
		windowRepo:             windowRepo,
		windowDishRelationRepo: windowDishRelationRepo,
		windowCommentRepo:      windowCommentRepo,
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

	if dto.Rate < 1 || dto.Rate > 5 {
		return 0, fmt.Errorf("rate must be between 1 and 5")
	}

	return cs.canteenCommentRepo.AddCanteenComment(dto, userID)
}

func (cs *CommentService) AddWindowComment(dto *dto.AddWindowCommentRequest, userID uint) (uint, error) {
	exists, err := cs.windowRepo.CheckWindowExists(dto.WindowID)
	if err != nil {
		return 0, err
	}
	if !exists {
		return 0, fmt.Errorf("window with id %d does not exist", dto.WindowID)
	}

	if dto.Rate < 1 || dto.Rate > 5 {
		return 0, fmt.Errorf("rate must be between 1 and 5")
	}

	if dto.DishID != nil {
		dishInWindow, err := cs.windowDishRelationRepo.CheckDishInWindow(dto.WindowID, *dto.DishID)
		if err != nil {
			return 0, err
		}
		if !dishInWindow {
			return 0, fmt.Errorf("dish with id %d is notin window %d", *dto.DishID, dto.WindowID)
		}
	}

	return cs.windowCommentRepo.AddWindowComment(dto, userID)
}
