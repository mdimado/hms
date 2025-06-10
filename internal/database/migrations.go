package database

import (
	"gorm.io/gorm"
	"hospital-management/internal/models"
)

func RunMigrations(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.User{},
		&models.Patient{},
	)
}
