package entity

type Tax struct {
	TaxID       string  `json:"tax_id" gorm:"primaryKey"`
	TaxRate     float64 `json:"tax_rate"`
	TaxType     string  `json:"tax_type"`
	Description string  `json:"description"`
}
