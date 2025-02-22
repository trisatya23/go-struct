package database

import (
	"go-struct/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "root:admin@tcp(localhost:3306)/struct?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	database.AutoMigrate(
		&entity.Product{},
		&entity.Customer{},
		&entity.Discount{},
		&entity.Employee{},
		&entity.Inventory{},
		&entity.Order{},
		&entity.OrderItem{},
		&entity.Payment{},
		&entity.Receipt{},
		&entity.Tax{},
	)
	DB = database
}
