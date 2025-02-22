package entity

type OrderItem struct {
	ID         uint    `json:"id" gorm:"primaryKey;autoIncrement"`
	OrderID    string  `json:"order_id" gorm:"index"`
	ProductID  string  `json:"product_id"`
	Quantity   int     `json:"quantity"`
	UnitPrice  float64 `json:"unit_price"`
	TotalPrice float64 `json:"total_price"`
}
