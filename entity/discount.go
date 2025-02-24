package entity

import "time"

type Discount struct {
	DiscountID  string    `json:"discount_id" gorm:"type:char(36);primaryKey"`
	Description string    `json:"description" gorm:"not null"`
	DiscountPct float64   `json:"discount_pct" gorm:"not null"`
	ValidFrom   time.Time `json:"valid_from" gorm:"type:datetime;not null"`
	ValidUntil  time.Time `json:"valid_until" gorm:"type:datetime;not null"`
}
