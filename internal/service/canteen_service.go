package service

import (
	dto_b "SJTU-Canteen-Community/internal/dto/b"
	dto_c "SJTU-Canteen-Community/internal/dto/c"
	"SJTU-Canteen-Community/internal/model"
	"SJTU-Canteen-Community/internal/repository"
	"fmt"

	log "github.com/sirupsen/logrus"
)

type CanteenService struct {
	canteenRepo            *repository.CanteenRepository
	windowRepo             *repository.WindowRepository
	windowDishRelationRepo *repository.WindowDishRelationRepository
	dishRepo               *repository.DishRepository
}

func NewCanteenService(canteenRepo *repository.CanteenRepository, windowRepo *repository.WindowRepository, windowDishRelationRepo *repository.WindowDishRelationRepository, dishRepo *repository.DishRepository) *CanteenService {
	return &CanteenService{canteenRepo: canteenRepo, windowRepo: windowRepo, windowDishRelationRepo: windowDishRelationRepo, dishRepo: dishRepo}
}

func (s *CanteenService) AddCanteen(dto *dto_b.AddCanteenRequest) (uint, error) {
	return s.canteenRepo.AddCanteen(dto)
}

func (s *CanteenService) MAddWindowsToCanteen(dto *dto_b.MAddWindowsToCanteenRequest) ([]model.Window, error) {
	exists, err := s.canteenRepo.CheckCanteenExists(dto.CanteenID)
	if err != nil {
		log.Errorf("Failed to check canteen existence: %v", err)
		return nil, err
	}
	if !exists {
		return nil, fmt.Errorf("Canteen with ID %d does not exist", dto.CanteenID)
	}

	return s.windowRepo.AddWindowsToCanteen(dto)
}

func (s *CanteenService) MAddDishesToWindow(dto *dto_b.MAddDishesToWindowRequest) error {
	windowExists, err := s.windowRepo.CheckWindowExists(dto.WindowID)
	if err != nil {
		log.Errorf("Failed to check window existence: %v", err)
		return err
	}

	if !windowExists {
		return fmt.Errorf("Window with ID %d does not exist", dto.WindowID)
	}

	dishIDs := make([]uint, len(dto.Dishes))
	for i, dish := range dto.Dishes {
		dishIDs[i] = dish.ID
	}
	allDishesExists, err := s.dishRepo.MCheckAllDishesExist(dishIDs)
	if err != nil {
		log.Errorf("Failed to check all dishes existence: %v", err)
		return err
	}

	if !allDishesExists {
		return fmt.Errorf("One or more dishes do not exist")
	}

	return s.windowDishRelationRepo.MAddDishesToWindow(dto)
}

func (s *CanteenService) ListCanteens(dto *dto_c.ListCanteensRequest) (dto_c.ListCanteensResponse, error) {
	return s.canteenRepo.ListCanteens(dto)
}

func (s *CanteenService) ListWindowsOfCanteen(dto *dto_c.ListWindowsOfCanteenRequest) (dto_c.ListWindowsOfCanteenResponse, error) {
	return s.windowRepo.ListWindowsOfCanteen(dto)
}

func (s *CanteenService) ListDishesOfWindow(dto *dto_c.ListDishesOfWindowRequest) (dto_c.ListDishesOfWindowResponse, error) {
	dishIDs, total, err := s.windowDishRelationRepo.ListDishIDsOfWindow(*dto)
	if err != nil {
		return dto_c.ListDishesOfWindowResponse{}, err
	}
	dishes, err := s.dishRepo.ListDishesByIDs(dishIDs)
	if err != nil {
		return dto_c.ListDishesOfWindowResponse{}, err
	}
	return dto_c.ListDishesOfWindowResponse{
		Dishes: dishes,
		Total:  total,
	}, nil
}
