package database

import (
	"errors"

	"github.com/jinzhu/gorm"

	"sample/app/domain"
)

type UserRepository struct{}

func (repo *UserRepository) Find(db *gorm.DB) (users []domain.Users, err error) {
	users = []domain.Users{}
	db.Find(&users)
	return users, nil
}

func (repo *UserRepository) FindByID(db *gorm.DB, id int) (user domain.Users, err error) {
	user = domain.Users{}
	db.First(&user, id)
	if user.ID <= 0 {
		return domain.Users{}, errors.New("user is not found")
	}
	return user, nil
}

func (repo *UserRepository) Create(db *gorm.DB, user *domain.Users) (err error) {
	result := db.Create(&user)
	return result.Error
}
