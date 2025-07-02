package config

import (
	"log"

	"BuhPro+/internal/domain" // Обновите импорт на buhpro/internal/domain

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(url string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err) //
	}

	err = db.AutoMigrate(
		&domain.User{},
		&domain.RefreshToken{},
		&domain.Customer{},
		&domain.Coach{},
		&domain.Executor{},
	)
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err) // [cite: 4]
	}

	return db
}
