package usecase

import (
	"github.com/jinzhu/gorm"
)

type DBRepository interface {
	Begin() *gorm.DB
	Connect() *gorm.DB
}
