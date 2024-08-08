package main

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Model User
type User struct {
	ID        uint    `gorm:"primaryKey;autoIncrement"`
	Name      string  `gorm:"type:varchar(100);not null"`
	Email     string  `gorm:"type:varchar(100);uniqueIndex;not null"`
	Profile   Profile `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Model Profile
type Profile struct {
	ID        uint   `gorm:"primaryKey;autoIncrement"`
	UserID    uint   `gorm:"unique;not null"`
	Age       int    `gorm:"not null"`
	Address   string `gorm:"type:varchar(255);not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func main() {
	// DSN untuk PostgreSQL
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Jakarta"

	// Membuka koneksi ke database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// Migrasi schema
	db.AutoMigrate(&User{}, &Profile{})

	// Membuat user dan profile baru
	user := User{
		Name:  "Budi Santoso",
		Email: "budi.santoso@example.com",
		Profile: Profile{
			Age:     30,
			Address: "Jl. Raya No. 1",
		},
	}

	// Menyimpan user dan profile baru ke database
	if err := db.Create(&user).Error; err != nil {
		log.Fatalf("failed to create user: %v", err)
	}

	GetDataByID(db, 1)
}

func GetDataByID(db *gorm.DB, id int) {
	// Mengambil user beserta profile-nya
	var fetchedUser User
	if err := db.Preload("Profile").First(&fetchedUser, id).Error; err != nil {
		log.Fatalf("failed to fetch user: %v", err)
	}

	fmt.Printf("User: %+v\n", fetchedUser)
	fmt.Printf("Profile: %+v\n", fetchedUser.Profile)
}

func AllDataUser(db *gorm.DB) {
	// Mengambil data menggunakan joins
	var results []struct {
		UserID         uint
		UserName       string
		UserEmail      string
		ProfileAge     int
		ProfileAddress string
	}

	err := db.Table("users").
		Joins("JOIN profiles ON profiles.user_id = users.id").
		Select("users.id as user_id, users.name as user_name, users.email as user_email, profiles.age as profile_age, profiles.address as profile_address").
		Scan(&results).Error

	if err != nil {
		log.Fatalf("failed to query with joins: %v", err)
	}

	for _, result := range results {
		fmt.Printf("User ID: %d, Name: %s, Email: %s, Age: %d, Address: %s\n",
			result.UserID, result.UserName, result.UserEmail, result.ProfileAge, result.ProfileAddress)
	}
}

func UpsertDataUser(db *gorm.DB) {
	// Menyimpan atau memperbarui data User dan Profile
	// Misalkan kita ingin meng-upsert user dengan ID 1
	upsertUser := User{
		ID:    1,
		Name:  "Budi Santoso",
		Email: "budi.santoso@example.com",
		Profile: Profile{
			Age:     31,
			Address: "Jl. Raya No. 2",
		},
	}

	// Auto Upsert: Simpan atau perbarui User dan Profile
	if err := db.Save(&upsertUser).Error; err != nil {
		log.Fatalf("failed to upsert user: %v", err)
	}
}

func SkipUpsertDataUser(db *gorm.DB) {
	// mengupdate user tanpa mengubah profile-nya
	user := User{
		ID:   1,
		Name: "Budi Prasetyo",
	}

	// Menggunakan Omit untuk mengabaikan update pada Profile
	if err := db.Omit("Profile").Save(&user).Error; err != nil {
		log.Fatalf("failed to update user: %v", err)
	}

}
