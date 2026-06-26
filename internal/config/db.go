package config

import (
	"fmt"
	"myapp/internal/database"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase(config *Config) *gorm.DB {
	dsn := config.DatabaseURL

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		TranslateError: true,
	})

	if err != nil {
		fmt.Println("server error:", err)
		return nil
	}

	database.Migrate(db)
	fmt.Println("Server connected successfully!")
	
	return db
}