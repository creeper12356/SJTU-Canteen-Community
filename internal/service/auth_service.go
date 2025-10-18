package service

import (
	dto "SJTU-Canteen-Community/internal/dto/c"
	"SJTU-Canteen-Community/internal/repository"

	log "github.com/sirupsen/logrus"
)

type AuthService struct {
	UserRepo *repository.UserRepository
}

func NewAuthService(UserRepo *repository.UserRepository) *AuthService {
	return &AuthService{UserRepo: UserRepo}
}

func (as *AuthService) Login(loginReq dto.LoginRequest) (uint, error) {
	user, err := as.UserRepo.FindUserByEmail(loginReq.Email)
	if err != nil {
		log.Errorf("Login failed: %v", err)
		return 0, err
	}

	return user.ID, nil
}
