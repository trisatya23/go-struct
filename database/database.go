package database

import (
	"go-struct/entity"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() *gorm.DB {
	dsn := "root:admin@tcp(localhost:3306)/struct?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("Database connected successfully!")

	// Migrasi database dengan urutan yang tepat
	err = database.Debug().AutoMigrate(
		&entity.Customer{},
		&entity.Employee{},
		&entity.Product{},
		&entity.Inventory{},
		&entity.Discount{},
		&entity.Order{},
		&entity.OrderItem{},
		&entity.Payment{},
		&entity.Receipt{},
		&entity.Tax{},
	)

	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// Set variabel global DB
	DB = database
	return DB
}
