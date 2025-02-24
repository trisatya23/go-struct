package entity

type OrderItem struct {
	OrderItemID uint    `json:"order_item_id" gorm:"primaryKey;autoIncrement"`
	OrderID     uint    `gorm:"index"`
	ProductID   string  `json:"product_id"`
	Quantity    int     `json:"quantity"`
	UnitPrice   float64 `json:"unit_price"`
	TotalPrice  float64 `json:"total_price"`
}
