package Orders

import "./Person"

// Orders Model
type Orders struct {
	OrderID     int
	OrderNumber int
	Person
}
