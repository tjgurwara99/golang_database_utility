package entity

import "time"

// Supplier  Struct to represent Supplier of a Company
type Supplier struct {
	SupplierID  int64
	LastContact time.Time
	*Profile
	*Company
}
