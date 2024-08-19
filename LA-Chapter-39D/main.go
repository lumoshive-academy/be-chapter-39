// package main

// import (
// 	"fmt"
// 	"log"
// 	"time"

// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// 	"gorm.io/gorm/logger"
// )

// // Model User
// type User struct {
// 	ID        uint    `gorm:"primaryKey;autoIncrement"`
// 	Name      string  `gorm:"type:varchar(100);not null"`
// 	Email     string  `gorm:"type:varchar(100);uniqueIndex;not null"`
// 	Orders    []Order `gorm:"foreignKey:UserID"`
// 	CreatedAt time.Time
// 	UpdatedAt time.Time
// }

// // Model Order
// type Order struct {
// 	ID        uint    `gorm:"primaryKey;autoIncrement"`
// 	OrderID   string  `gorm:"type:varchar(50);not null"`
// 	UserID    uint    `gorm:"not null"`
// 	User      User    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
// 	Amount    float64 `gorm:"type:decimal(10,2);not null"`
// 	CreatedAt time.Time
// 	UpdatedAt time.Time
// }

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

// 	// Migrasi schema
// 	db.AutoMigrate(&User{}, &Order{})

// 	// Membuat user baru
// 	user := User{
// 		Name:  "Budi Santoso",
// 		Email: "budi.santoso@example.com",
// 		Orders: []Order{
// 			{OrderID: "INV001", Amount: 100000},
// 			{OrderID: "INV002", Amount: 250000},
// 		},
// 	}

// 	// Menyimpan user beserta order ke database
// 	if err := db.Create(&user).Error; err != nil {
// 		log.Fatalf("failed to create user: %v", err)
// 	}

// 	// implementation on to many
// 	OnToManyGrom(db, int(user.ID))

// 	// implementation belongs to
// 	BelongToGorm(db, 1)

// }

// func OnToManyGrom(db *gorm.DB, id int) {

// 	// Mengambil user beserta order-nya
// 	var fetchedUser User
// 	if err := db.Preload("Orders").First(&fetchedUser, id).Error; err != nil {
// 		log.Fatalf("failed to fetch user: %v", err)
// 	}

// 	fmt.Printf("User: %+v\n", fetchedUser)
// 	fmt.Printf("Orders: %+v\n", fetchedUser.Orders)
// }

// func BelongToGorm(db *gorm.DB, id int) {
// 	// Mengambil order beserta user-nya (Belongs To)
// 	var fetchedOrder Order
// 	if err := db.Preload("User").First(&fetchedOrder, id).Error; err != nil {
// 		log.Fatalf("failed to fetch order: %v", err)
// 	}

// 	fmt.Printf("Order: %+v\n", fetchedOrder)
// 	fmt.Printf("User: %+v\n", fetchedOrder.User)
// }

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
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	Name      string    `gorm:"type:varchar(100);not null"`
	Email     string    `gorm:"type:varchar(100);uniqueIndex;not null"`
	Projects  []Project `gorm:"many2many:user_projects;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Model Project
type Project struct {
	ID        uint   `gorm:"primaryKey;autoIncrement"`
	Name      string `gorm:"type:varchar(100);not null"`
	Users     []User `gorm:"many2many:user_projects;"`
	CreatedAt time.Time
	UpdatedAt time.Time
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
	db.AutoMigrate(&User{}, &Project{})

	// Membuat user dan project baru
	user1 := User{Name: "Budi Santoso", Email: "budi.santoso@example.com"}
	user2 := User{Name: "Siti Aminah", Email: "siti.aminah@example.com"}

	project1 := Project{Name: "Project A"}
	project2 := Project{Name: "Project B"}

	// Menambahkan relasi many-to-many
	db.Model(&user1).Association("Projects").Append(&project1, &project2)
	db.Model(&user2).Association("Projects").Append(&project1)

	// Menyimpan user dan project baru ke database
	db.Save(&user1)
	db.Save(&user2)

	// Mengambil user beserta projects-nya
	var fetchedUser User
	if err := db.Preload("Projects").First(&fetchedUser, user1.ID).Error; err != nil {
		log.Fatalf("failed to fetch user: %v", err)
	}

	fmt.Printf("User: %+v\n", fetchedUser)
	fmt.Println("Projects:")
	for _, project := range fetchedUser.Projects {
		fmt.Printf("- %+v\n", project)
	}
}
