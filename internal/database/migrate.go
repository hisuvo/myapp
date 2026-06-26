package database

import (
	"myapp/internal/event"
	"myapp/internal/users"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
    db.AutoMigrate(
        &event.Event{},
        &users.User{},
    )
}