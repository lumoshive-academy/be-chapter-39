package main

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	ID        uint   `gorm:"primaryKey;autoIncrement"`
	Name      string `gorm:"type:varchar(100);not null"`
	Email     string `gorm:"type:varchar(100);uniqueIndex;not null"`
	Password  string `gorm:"type:varchar(255);not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`
}

func main() {
	dsn := "host=localhost user=postgres password=mysecretpassword dbname=testdb port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// Migrasi model ke database
	db.AutoMigrate(&User{})

	// Membuat user baru
	users := []User{
		{Name: "Budi Santoso", Email: "budi.santoso@example.com", Password: "password"},
		{Name: "Siti Nurhaliza", Email: "siti.nurhaliza@example.com", Password: "password"},
		{Name: "Agus Salim", Email: "agus.salim@example.com", Password: "password"},
	}
	db.Create(&users)

}

func DeleteUser(db *gorm.DB) {

	// Menghapus user berdasarkan primary key (ID)
	db.Delete(&User{}, 1)

	fmt.Printf("User with ID %d has been deleted\n", 1)

	// Menghapus user dengan kondisi (menghapus berdasarkan email)
	db.Delete(&User{}, "email = ?", "siti.nurhaliza@example.com")

	fmt.Println("User with email siti.nurhaliza@example.com has been deleted")
}
