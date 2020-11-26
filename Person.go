package Person

import "./Orders"

// Person Model
type Person struct {
	PersonID int
	Name     string
	Orders   []Orders
}
