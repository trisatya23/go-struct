package entity

import "time"

type Payment struct {
	PaymentID   string    `json:"payment_id" gorm:"type:char(36);primaryKey"`
	OrderID     uint      `json:"order_id" gorm:"not null;index"`
	Amount      float64   `json:"amount" gorm:"type:decimal(10,2);not null"`
	PaymentType string    `json:"payment_type" gorm:"type:varchar(50);not null"`
	PaymentDate time.Time `json:"payment_date" gorm:"type:datetime;not null"`
	Status      string    `json:"status" gorm:"type:varchar(20);not null"`
}
