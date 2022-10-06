package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	log.Printf("Gorm tutorial")

	// init db
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	//var p Product
	//db.First(&p, "price = ?", 1000)
	//log.Println("p:", p)
	//os.Exit(0)

	// Migrate schema
	db.AutoMigrate(&Product{})

	// CRUD
	// Create
	db.Create(&Product{Code: "A1", Price: 100})
	db.Create(&Product{Code: "B1", Price: 50})
	db.Create(&Product{Code: "C1", Price: 1000})

	// Read
	var product1, product2, product3 Product
	db.First(&product1, 1)
	db.First(&product2, "code = ?", "B1")
	db.First(&product3, "price = ?", 1000)
	log.Println("Product with primary key '1'", product1)
	log.Println("Product with code 'D42':", product2)
	log.Println("Product with price 1000:", product3)

	// Update
	db.Model(&product1).Update("Price", 200)
	db.Model(&product2).Updates(Product{Price: 100, Code: "B2"}) // update multiple values on same object

	// Delete
	db.Delete(&product3, 1)

}
