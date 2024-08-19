// package main

// import (
// 	"fmt"
// 	"log"
// 	"time"

// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// 	"gorm.io/gorm/logger"
// )

// // Model Author
// type Author struct {
// 	ID        uint      `gorm:"primaryKey;autoIncrement"`
// 	Name      string    `gorm:"type:varchar(100);not null"`
// 	Books     []Book    `gorm:"foreignKey:AuthorID"` // relasi satu-ke-banyak
// 	CreatedAt time.Time
// 	UpdatedAt time.Time
// }

// // Model Book
// type Book struct {
// 	ID        uint      `gorm:"primaryKey;autoIncrement"`
// 	Title     string    `gorm:"type:varchar(100);not null"`
// 	AuthorID  uint      `gorm:"not null"`
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
// 	db.AutoMigrate(&Author{}, &Book{})

// 	// Membuat author dan book baru
// 	author := Author{Name: "J.K. Rowling"}
// 	book1 := Book{Title: "Harry Potter and the Philosopher's Stone"}
// 	book2 := Book{Title: "Harry Potter and the Chamber of Secrets"}

// 	// Menambahkan books ke author
// 	db.Model(&author).Association("Books").Append(&book1, &book2)

// 	// Menyimpan author dan books ke database
// 	db.Save(&author)
// 	db.Save(&book1)
// 	db.Save(&book2)

// 	// Mengambil author beserta books-nya
// 	var fetchedAuthor Author
// 	if err := db.Preload("Books").First(&fetchedAuthor, author.ID).Error; err != nil {
// 		log.Fatalf("failed to fetch author: %v", err)
// 	}

// 	fmt.Printf("Author: %+v\n", fetchedAuthor)
// 	fmt.Println("Books:")
// 	for _, book := range fetchedAuthor.Books {
// 		fmt.Printf("- %+v\n", book)
// 	}
// }

package main

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
)

// Model Author
type Author struct {
	ID        uint   `gorm:"primaryKey;autoIncrement"`
	Name      string `gorm:"type:varchar(100);not null"`
	Books     []Book `gorm:"foreignKey:AuthorID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Model Book
type Book struct {
	ID        uint     `gorm:"primaryKey;autoIncrement"`
	Title     string   `gorm:"type:varchar(100);not null"`
	AuthorID  uint     `gorm:"not null"`
	Reviews   []Review `gorm:"foreignKey:BookID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Model Review
type Review struct {
	ID        uint   `gorm:"primaryKey;autoIncrement"`
	Content   string `gorm:"type:text"`
	BookID    uint   `gorm:"not null"`
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
	db.AutoMigrate(&Author{}, &Book{}, &Review{})

	// Membuat data contoh
	author := Author{Name: "J.K. Rowling"}
	book1 := Book{Title: "Harry Potter and the Philosopher's Stone"}
	book2 := Book{Title: "Harry Potter and the Chamber of Secrets"}
	review1 := Review{Content: "Amazing book!"}
	review2 := Review{Content: "A must-read for fantasy lovers."}

	// Menambahkan relasi
	book1.Reviews = []Review{review1}
	book2.Reviews = []Review{review2}
	author.Books = []Book{book1, book2}

	// Menyimpan data ke database
	db.Save(&author)
	db.Save(&book1)
	db.Save(&book2)
	db.Save(&review1)
	db.Save(&review2)

	// Contoh Preloading
	var authors []Author

	// Preloading dengan kondisi
	fmt.Println("Preloading dengan kondisi (hanya buku yang memiliki ulasan)")
	if err := db.Preload("Books", "id IN (?)", []uint{book1.ID}).Preload("Books.Reviews").Find(&authors).Error; err != nil {
		log.Fatalf("failed to fetch authors: %v", err)
	}
	for _, author := range authors {
		fmt.Printf("Author: %+v\n", author)
		for _, book := range author.Books {
			fmt.Printf("Book: %+v\n", book)
			for _, review := range book.Reviews {
				fmt.Printf("Review: %+v\n", review)
			}
		}
	}

	// Preloading dengan nested struct
	fmt.Println("Preloading dengan nested struct")
	if err := db.Preload("Books.Reviews").Find(&authors).Error; err != nil {
		log.Fatalf("failed to fetch authors: %v", err)
	}
	for _, author := range authors {
		fmt.Printf("Author: %+v\n", author)
		for _, book := range author.Books {
			fmt.Printf("Book: %+v\n", book)
			for _, review := range book.Reviews {
				fmt.Printf("Review: %+v\n", review)
			}
		}
	}

	// Contoh Preloading dengan clauses.Associations
	fmt.Println("Preloading dengan clauses.Associations")
	if err := db.Preload(clause.Associations).Find(&authors).Error; err != nil {
		log.Fatalf("failed to fetch authors: %v", err)
	}

	for _, author := range authors {
		fmt.Printf("Author: %+v\n", author)
		for _, book := range author.Books {
			fmt.Printf("Book: %+v\n", book)
			for _, review := range book.Reviews {
				fmt.Printf("Review: %+v\n", review)
			}
		}
	}
}
