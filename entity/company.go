package entity

import "time"

// Company Entity
type Company struct {
	CompanyID       uint64    `json:"company_id" db:"compant_id"`
	CompanyName     string    `json:"company_name" db:"company_name"`
	CompanyIsActive bool      `json:"company_is_active" db:"is_active"`
	LastPayment     time.Time `json:"last_payment" db:"last_payment"`
	MainPosition    string    `json:"main_position" db:"main_position"`
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
