package models

// Orders model
type Orders struct {
	OrderID     int
	OrderNumber int
	*Person
}

// NewOrder constructor
func NewOrder(orderID, orderNumber int, person *Person) Orders {
	order := Orders{
		OrderID:     orderID,
		OrderNumber: orderNumber,
		Person:      person,
	}
	return order
}
