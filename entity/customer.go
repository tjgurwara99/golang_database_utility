package entity

import "time"

// Customer Struct to represent Customers of a Company
type Customer struct {
	CustomerID  int64
	LastContact time.Time
	*Profile
	*Company
}
