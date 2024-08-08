package main

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
)

// Menggunakan gorm.Model sebagai embedded struct
type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(100);not null"`
	Email    string `gorm:"type:varchar(100);uniqueIndex;not null"`
	Password string `gorm:"type:varchar(255);not null"`
}

func main() {
	// DSN untuk PostgreSQL
	dsn := "host=localhost user=postgres password=mysecretpassword dbname=testdb port=5432 sslmode=disable TimeZone=Asia/Jakarta"

	// Membuka koneksi ke database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// Migrasi schema
	db.AutoMigrate(&User{})

	// Membuat user baru
	user := User{
		Name:     "Budi Santoso",
		Email:    "budi.santoso@example.com",
		Password: "password",
	}

	// Menyimpan user baru ke database
	db.Create(&user)

	fmt.Printf("User created with ID: %d\n", user.ID)
	fmt.Printf("CreatedAt: %s\n", user.CreatedAt)
	fmt.Printf("UpdatedAt: %s\n", user.UpdatedAt)
}

func LockDataWhenUpdate(db *gorm.DB) {
	// Simulasi update dengan locking
	db.Transaction(func(tx *gorm.DB) error {
		var user User

		// Mengunci record user dengan ID 1 untuk update
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&user, 1).Error; err != nil {
			return err
		}

		// Update nama user
		user.Name = "Budi Santoso Updated"
		if err := tx.Save(&user).Error; err != nil {
			return err
		}

		fmt.Printf("User with ID %d has been updated to %s\n", user.ID, user.Name)
		return nil
	})
}
