package repositories

import (
	"dans-multi-pro/models"

	"gorm.io/gorm"
)

type UserRepo interface {
	CreateUser(user *models.User) (*models.User, error)
	CheckUserByUsername(username string, user *models.User) error
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return &userRepo{db}
}

func (u *userRepo) CreateUser(user *models.User) (*models.User, error) {
	return user, u.db.Create(&user).Error
}

func (u *userRepo) CheckUserByUsername(username string, user *models.User) error {
	return u.db.Where("username=?", username).Take(&user).Error
}
