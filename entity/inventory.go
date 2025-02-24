package entity

import "time"

type Inventory struct {
	ProductID    string    `json:"product_id" gorm:"type:char(36);not null"`
	StockQty     int       `json:"stock_qty" gorm:"not null"`
	RestockLevel int       `json:"restock_level" gorm:"not null"`
	LastRestock  time.Time `json:"last_restock" gorm:"type:datetime;not null"`
}
