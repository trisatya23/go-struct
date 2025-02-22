package entity

type Discount struct {
	DiscountID  string  `json:"discount_id" gorm:"primaryKey"`
	Description string  `json:"description"`
	DiscountPct float64 `json:"discount_pct"`
	ValidFrom   string  `json:"valid_from"`
	ValidUntil  string  `json:"valid_until"`
}
