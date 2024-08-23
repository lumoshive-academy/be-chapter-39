package database

import (
	"golang-chapter-39/LA-Chapter-39H/models"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
}
