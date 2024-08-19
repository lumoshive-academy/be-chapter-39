// package main

// import (
// 	"log"
// 	"time"

// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// 	"gorm.io/gorm/logger"
// )

// func main() {
// 	// DSN untuk PostgreSQL
// 	dsn := "host=localhost user=postgres password=mysecretpassword dbname=testdb port=5432 sslmode=disable TimeZone=Asia/Jakarta"

// 	// Membuka koneksi ke database
// 	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
// 		Logger: logger.Default.LogMode(logger.Info),
// 	})
// 	if err != nil {
// 		log.Fatalf("failed to connect database: %v", err)
// 	}

// 	// Mengambil objek *sql.DB dari gorm.DB
// 	sqlDB, err := db.DB()
// 	if err != nil {
// 		log.Fatalf("failed to get sql.DB: %v", err)
// 	}

// 	// Mengatur connection pool
// 	sqlDB.SetMaxOpenConns(10)                  // Jumlah maksimum koneksi terbuka
// 	sqlDB.SetMaxIdleConns(5)                   // Jumlah maksimum koneksi idle
// 	sqlDB.SetConnMaxLifetime(30 * time.Minute) // Lama waktu koneksi
// 	sqlDB.SetConnMaxIdleTime(5 * time.Minute)  // Lama waktu idle koneksi

// }

package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Model User
type User struct {
	ID        uint   `gorm:"primaryKey;autoIncrement"`
	Name      string `gorm:"type:varchar(100);not null"`
	Email     string `gorm:"type:varchar(100);uniqueIndex;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Before Create hook
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	// Mengubah nama menjadi huruf kapital sebelum menyimpan
	u.Name = strings.ToUpper(u.Name)
	return nil
}

// After Create hook
func (u *User) AfterCreate(tx *gorm.DB) (err error) {
	// Mencatat waktu pembuatan pengguna
	fmt.Printf("User created at: %v\n", u.CreatedAt)
	return nil
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
	if err := db.AutoMigrate(&User{}); err != nil {
		log.Fatalf("failed to migrate schema: %v", err)
	}

	// Membuat user baru
	user := User{
		Name:  "budi santoso",
		Email: "budi.santoso@example.com",
	}

	if err := db.Create(&user).Error; err != nil {
		log.Fatalf("failed to create user: %v", err)
	}

	// Mengambil data untuk memastikan hooks bekerja
	var fetchedUser User
	if err := db.First(&fetchedUser, user.ID).Error; err != nil {
		log.Fatalf("failed to fetch user: %v", err)
	}
	fmt.Printf("Fetched User: %+v\n", fetchedUser)
}
