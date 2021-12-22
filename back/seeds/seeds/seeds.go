package seeds

import (
	"time"

	"github.com/jinzhu/gorm"
)

func All() []Seed {
	return []Seed{
		{
			Name: "CreateJane",
			Run: func(db *gorm.DB) error {
				return CreateUser(db, "Jane", "pass", "jane@example.com", time.Now(), "https://avatars.githubusercontent.com/u/49118806")
			},
		},
		{
			Name: "CreateJohn",
			Run: func(db *gorm.DB) error {
				return CreateUser(db, "John", "pass", "john@example.com", time.Now(), "https://avatars.githubusercontent.com/u/49118806")
			},
		},
	}
}
