package repository

import (
	"errors"

	"github.com/mardini1414/simple-procurement-system-be/internal/dto"
	"github.com/mardini1414/simple-procurement-system-be/internal/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Create(req dto.RegisterRequest) (*model.User, error) {
	user := model.User{
		Username: req.Username,
		Password: req.Password,
		Role:     "user",
	}

	if err := r.db.Create(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) FindByUsername(username string) (*model.User, error) {
	var user model.User
	err := r.db.Take(&user, "username = ?", username).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	return &user, nil
}
