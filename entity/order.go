package entity

import "time"

type Order struct {
	OrderID     uint        `json:"order_id" gorm:"primaryKey;autoIncrement"`
	CustomerID  string      `json:"customer_id" gorm:"type:char(36);not null"`
	OrderDate   time.Time   `json:"order_date" gorm:"type:datetime;not null"`
	TotalAmount float64     `json:"total_amount" gorm:"type:decimal(10,2);not null"`
	OrderItems  []OrderItem `json:"order_items" gorm:"foreignKey:OrderID"`
}
