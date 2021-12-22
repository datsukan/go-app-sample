package seeds

import (
	"time"

	"portfolio/app/domain"

	"github.com/jinzhu/gorm"
)

func CreateUser(db *gorm.DB, username string, password string, email string, birthday time.Time, image_url string) error {
	return db.Create(&domain.Users{Username: username, Password: password, Email: email, Birthday: birthday, ImageUrl: image_url}).Error
}
