package repositories

import (
	"gorm.io/gorm"
)

type JobRepo interface {
}

type jobRepo struct {
	db *gorm.DB
}

func NewJobRepo(db *gorm.DB) UserRepo {
	return &userRepo{db}
}
