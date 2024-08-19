// package main

// import (
// 	"fmt"
// 	"log"
// 	"time"

// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// 	"gorm.io/gorm/logger"
// )

// // Model Book
// type Book struct {
// 	ID        uint    `gorm:"primaryKey;autoIncrement"`
// 	Title     string  `gorm:"type:varchar(100);not null"`
// 	AuthorID  uint    `gorm:"not null"`
// 	Price     float64 `gorm:"type:decimal(10,2);not null"`
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
// 	db.AutoMigrate(&Book{})

// 	// Menambahkan data contoh
// 	books := []Book{
// 		{Title: "Book 1", AuthorID: 1, Price: 29.99},
// 		{Title: "Book 2", AuthorID: 1, Price: 39.99},
// 		{Title: "Book 3", AuthorID: 2, Price: 49.99},
// 	}
// 	db.Create(&books)

// 	// Contoh Query Aggregation
// 	var count int64
// 	if err := db.Model(&Book{}).Where("author_id = ?", 1).Count(&count).Error; err != nil {
// 		log.Fatalf("failed to count books: %v", err)
// 	}
// 	fmt.Printf("Number of books by author 1: %d\n", count)

// 	var totalPrice float64
// 	if err := db.Model(&Book{}).Where("author_id = ?", 1).Select("SUM(price)").Scan(&totalPrice).Error; err != nil {
// 		log.Fatalf("failed to sum book prices: %v", err)
// 	}
// 	fmt.Printf("Total price of books by author 1: %.2f\n", totalPrice)

// 	var averagePrice float64
// 	if err := db.Model(&Book{}).Where("author_id = ?", 1).Select("AVG(price)").Scan(&averagePrice).Error; err != nil {
// 		log.Fatalf("failed to average book prices: %v", err)
// 	}
// 	fmt.Printf("Average price of books by author 1: %.2f\n", averagePrice)

// 	var minPrice, maxPrice float64
// 	if err := db.Model(&Book{}).Where("author_id = ?", 1).Select("MIN(price)").Scan(&minPrice).Error; err != nil {
// 		log.Fatalf("failed to find min book price: %v", err)
// 	}
// 	if err := db.Model(&Book{}).Where("author_id = ?", 1).Select("MAX(price)").Scan(&maxPrice).Error; err != nil {
// 		log.Fatalf("failed to find max book price: %v", err)
// 	}
// 	fmt.Printf("Min price of books by author 1: %.2f\n", minPrice)
// 	fmt.Printf("Max price of books by author 1: %.2f\n", maxPrice)

// 	type Result struct {
// 		AuthorID uint
// 		Count    int64
// 	}
// 	var results []Result
// 	if err := db.Model(&Book{}).Select("author_id, COUNT(*) as count").Group("author_id").Scan(&results).Error; err != nil {
// 		log.Fatalf("failed to group books by author: %v", err)
// 	}
// 	for _, result := range results {
// 		fmt.Printf("Author ID: %d has %d books\n", result.AuthorID, result.Count)
// 	}

// 	if err := db.Model(&Book{}).
// 		Select("author_id, COUNT(*) as count").
// 		Group("author_id").
// 		Having("COUNT(*) > ?", 1). // Menggunakan HAVING untuk hanya menyertakan penulis dengan lebih dari satu buku
// 		Scan(&results).Error; err != nil {
// 		log.Fatalf("failed to group and filter books: %v", err)
// 	}

// 	for _, result := range results {
// 		fmt.Printf("Author ID: %d has %d books\n", result.AuthorID, result.Count)
// 	}
// }

// package main

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"time"

// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// 	"gorm.io/gorm/logger"
// )

// // Model User
// type User struct {
// 	ID        uint   `gorm:"primaryKey;autoIncrement"`
// 	Name      string `gorm:"type:varchar(100);not null"`
// 	Email     string `gorm:"type:varchar(100);uniqueIndex;not null"`
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
// 	db.AutoMigrate(&User{})

// 	// Menambahkan data contoh
// 	users := []User{
// 		{Name: "Alice", Email: "alice@example.com"},
// 		{Name: "Bob", Email: "bob@example.com"},
// 		{Name: "Charlie", Email: "charlie@example.com"},
// 	}
// 	db.Create(&users)

// 	// Menggunakan context dengan batas waktu
// 	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
// 	defer cancel()

// 	var fetchedUser User
// 	if err := db.WithContext(ctx).Where("name = ?", "Alice").First(&fetchedUser).Error; err != nil {
// 		if err == gorm.ErrRecordNotFound {
// 			fmt.Println("User not found")
// 		} else {
// 			log.Fatalf("failed to fetch user: %v", err)
// 		}
// 	} else {
// 		fmt.Printf("Fetched User: %+v\n", fetchedUser)
// 	}
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
	ID        uint   `gorm:"primaryKey;autoIncrement"`
	Name      string `gorm:"type:varchar(100);not null"`
	Email     string `gorm:"type:varchar(100);uniqueIndex;not null"`
	IsActive  bool   `gorm:"default:true"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Scope to filter active users
func ActiveUsers(db *gorm.DB) *gorm.DB {
	return db.Where("is_active = ?", true)
}

// Scope to filter users by name
func UsersByName(name string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("name = ?", name)
	}
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

	// Menambahkan data contoh
	users := []User{
		{Name: "Alice", Email: "alice@example.com", IsActive: true},
		{Name: "Bob", Email: "bob@example.com", IsActive: false},
		{Name: "Charlie", Email: "charlie@example.com", IsActive: true},
	}
	db.Create(&users)

	// Menggunakan scope untuk mendapatkan pengguna aktif
	var activeUsers []User
	if err := ActiveUsers(db).Find(&activeUsers).Error; err != nil {
		log.Fatalf("failed to fetch active users: %v", err)
	}
	fmt.Println("Active Users:")
	for _, user := range activeUsers {
		fmt.Printf("User: %+v\n", user)
	}

	// Menggunakan scope untuk mendapatkan pengguna berdasarkan nama
	var alice User
	if err := UsersByName("Alice")(db).First(&alice).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			fmt.Println("User not found")
		} else {
			log.Fatalf("failed to fetch user: %v", err)
		}
	} else {
		fmt.Printf("User by Name: %+v\n", alice)
	}
}
