package entity

import "time"

// Company Entity
type Company struct {
	CompanyID       int64
	CompanyName     string
	CompanyManager  *User
	CompanyIsActive bool
	LastPayment     time.Time
}

// String function to return default value for fmt.Printf like functions
func (company *Company) String() string {
	return company.CompanyName
}
