package entity

type Employee struct {
	EmployeeID string `json:"employee_id" gorm:"primaryKey"`
	Name       string `json:"name"`
	Role       string `json:"role"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	DateHired  string `json:"date_hired"`
}
