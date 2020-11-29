package entities

import "fmt"

// Order model
type Order struct {
	OrderID     int
	OrderNumber int
	Person      *Person
}

func (order *Order) String() string {
	return fmt.Sprintf("{OrderID: %d, OrderNumber: %d, Person: %v}", order.OrderID, order.OrderNumber, order.Person)
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

//SelectAllQuery string for the query
func (order Order) SelectAllQuery() string {
	return "select * from orders"
}
