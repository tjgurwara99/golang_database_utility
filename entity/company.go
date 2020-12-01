package entity

import "time"

// Company Entity
type Company struct {
	CompanyID       int64
	CompanyName     string
	CompanyIsActive bool
	LastPayment     time.Time
}

// String function to return default value for fmt.Printf like functions
func (company *Company) String() string {
	return company.CompanyName
}

// NewCompany Company Constructor
func NewCompany(companyName *string, companyIsActive *bool, lastPayment *time.Time) (*Company, error) {
	return &Company{
		CompanyName:     *companyName,
		CompanyIsActive: *companyIsActive,
		LastPayment:     *lastPayment,
	}, nil
}
