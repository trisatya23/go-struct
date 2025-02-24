package entity

import "time"

type Employee struct {
	EmployeeID string    `json:"employee_id" gorm:"type:char(36);primaryKey"`
	Name       string    `json:"name" gorm:"not null"`
	Role       string    `json:"role" gorm:"not null"`
	Email      string    `json:"email" gorm:"unique;not null"`
	Phone      string    `json:"phone" gorm:"unique;not null"`
	DateHired  time.Time `json:"date_hired" gorm:"type:datetime;not null"`
}
