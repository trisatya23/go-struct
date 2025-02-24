package entity

type Customer struct {
	CustomerID string `json:"customer_id" gorm:"type:char(36);primaryKey"`
	Name       string `json:"name" gorm:"not null"`
	Email      string `json:"email" gorm:"unique;not null"`
	Phone      string `json:"phone" gorm:"unique;not null"`
	Address    string `json:"address"`
	LoyaltyPts int64  `json:"loyalty_points"`
	CreatedAt  int64  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt  int64  `json:"updated_at" gorm:"autoUpdateTime"`
}
