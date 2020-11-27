package models

import "fmt"

// Order model
type Order struct {
	OrderID     int
	OrderNumber int
	*Person
}

func (order *Order) String() string {
	return fmt.Sprintf("OrderID: %d, OrderNumber: %d, Person: %v", order.OrderID, order.OrderNumber, order.Person)
}

// NewOrder constructor
func NewOrder(orderID, orderNumber int, person *Person) Order {
	order := Order{
		OrderID:     orderID,
		OrderNumber: orderNumber,
		Person:      person,
	}
	return order
}

func (o Order) SelectAllQuery() string {
	return "select * from orders"
}
