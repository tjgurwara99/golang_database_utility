package Person

import "./Orders"

type Person struct {
	personId int
	name     string
	orders   []Orders
}
