package entity

import "time"

// Product Struct to represent Products offered by a Company
type Product struct {
	ProductID     int64
	Name          string
	Stock         int64
	Reserved      bool
	Description   string
	Barcode       string
	SalePrice     float64
	CostPrice     float64
	TaxPercentage float64
	VatTaxable    bool
	AddedBy       *User
	UpdatedBy     *User
	Active        bool
	DateCreated   time.Time
	DateUpdated   time.Time
	*Company
	*Category
	*Supplier
}
