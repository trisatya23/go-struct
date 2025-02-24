package entity

type Product struct {
	ProductID   string  `json:"product_id" gorm:"type:char(36);primaryKey"`
	Name        string  `json:"name" gorm:"type:varchar(255);not null"`
	Description string  `json:"description" gorm:"type:text"`
	Price       float64 `json:"price" gorm:"type:decimal(10,2);not null"`
	StockQty    int     `json:"stock_qty" gorm:"not null;"`
	Category    string  `json:"category" gorm:"type:varchar(100);not null"`
	SKU         string  `json:"sku" gorm:"type:varchar(100);unique;not null"`
	TaxRate     float64 `json:"tax_rate" gorm:"type:decimal(5,2);not null"`
}
