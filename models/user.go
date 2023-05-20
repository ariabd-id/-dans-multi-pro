package models

import (
	"dans-multi-pro/helpers"
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	Username  string    `gorm:"not null;uniqueIndex" json:"username" valid:"required~ Your username is required"`
	Password  string    `gorm:"not null" json:"password" form:"password" valid:"required~Your full password is required"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

func (u *User) BeforeCreate(db *gorm.DB) error {
	_, errCreate := govalidator.ValidateStruct(u)
	u.Password = helpers.HashPassword(u.Password)
	return errCreate
}
