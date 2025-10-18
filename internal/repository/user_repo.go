package repository

import (
	"SJTU-Canteen-Community/internal/model"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (u *UserRepository) FindUserByEmail(email string) (model.User, error) {
	var user model.User
	result := u.db.Where("email = ?", email).First(&user)
	return user, result.Error
}
