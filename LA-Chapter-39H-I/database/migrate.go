package database

import (
	"golang-chapter-39/LA-Chapter-39H-I/models"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
}
