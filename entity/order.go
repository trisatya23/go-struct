package entity

import "time"

type Order struct {
	OrderID     string      `json:"order_id" gorm:"primary_key"`
	CustomerID  string      `json:"customer_id"`
	OrderDate   time.Time   `json:"order_date"`
	TotalAmount float64     `json:"total_amount"`
	OrderItems  []OrderItem `json:"order_items" gorm:"foreignKey:OrderID"`
}
