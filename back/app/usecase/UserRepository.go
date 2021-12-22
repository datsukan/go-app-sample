package usecase

import (
	"github.com/jinzhu/gorm"

	"portfolio/app/domain"
)

type UserRepository interface {
	Find(db *gorm.DB) (users []domain.Users, err error)
	FindByID(db *gorm.DB, id int) (user domain.Users, err error)
	Create(db *gorm.DB, user *domain.Users) (err error)
}
