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
	DeletedAt gorm.DeletedAt `gorm:"-"`
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

	DeleteUser(db)

	SoftDeleteUser(db)

	GetAllUserSoftDelete(db)

	RestoreDataUser(db, 1)
}

func DeleteUser(db *gorm.DB) {

	// Menghapus user berdasarkan primary key (ID)
	db.Delete(&User{}, 1)

	fmt.Printf("User with ID %d has been deleted\n", 1)

	// Menghapus user dengan kondisi (menghapus berdasarkan email)
	db.Delete(&User{}, "email = ?", "siti.nurhaliza@example.com")

	fmt.Println("User with email siti.nurhaliza@example.com has been deleted")
}

func SoftDeleteUser(db *gorm.DB) {
	// Menghapus user (soft delete)
	db.Delete(&User{}, 1)
	fmt.Printf("User with ID %d has been soft deleted\n", 1)
}

func GetAllUsersIncludingSoftDeleted(db *gorm.DB) {
	// Mengembalikan semua user termasuk yang soft deleted
	var allUsers []User
	db.Unscoped().Find(&allUsers)
	fmt.Println("All users (including soft deleted):", allUsers)

	// Mengembalikan semua user tanpa yang soft deleted
	var activeUsers []User
	db.Find(&activeUsers)
	fmt.Println("Active users:", activeUsers)
}

func GetAllUserSoftDelete(db *gorm.DB) {
	// Menampilkan list data user yang sudah dihapus
	var deletedUsers []User
	db.Unscoped().Where("deleted_at IS NOT NULL").Find(&deletedUsers)

	fmt.Println("List User yang Dihapus:")
	for _, user := range deletedUsers {
		fmt.Printf("ID: %d, Name: %s, Email: %s, DeletedAt: %v\n", user.ID, user.Name, user.Email, user.DeletedAt)
	}
}

func RestoreDataUser(db *gorm.DB, id int) {
	// Restore user yang sudah dihapus (misalnya dengan ID 1)
	var deletedUser User
	if err := db.Unscoped().First(&deletedUser, id).Error; err != nil {
		fmt.Println("User not found:", err)
		return
	}

	db.Model(&deletedUser).Update("DeletedAt", nil)

	// Verifikasi bahwa user sudah di-restore
	var restoredUser User
	if err := db.First(&restoredUser, 1).Error; err != nil {
		fmt.Println("Failed to restore user:", err)
		return
	}
	fmt.Println("Restored user:", restoredUser)
}
