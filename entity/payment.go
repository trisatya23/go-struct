package entity

type Payment struct {
	PaymentID   string  `json:"payment_id" gorm:"primaryKey"`
	OrderID     string  `json:"order_id"`
	Amount      float64 `json:"amount"`
	PaymentType string  `json:"payment_type"`
	PaymentDate string  `json:"payment_date"`
	Status      string  `json:"status"`
}
