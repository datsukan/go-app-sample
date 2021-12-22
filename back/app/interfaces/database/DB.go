package database

import (
	"github.com/jinzhu/gorm"
)

type DB interface {
	Begin() *gorm.DB
	Connect() *gorm.DB
}
