package model

import (
	"database/sql"

	"github.com/tjgurwara99/golang_database_utility/entity"
)

//OrderModel Order Entity Model
type OrderModel struct {
	DB *sql.DB
}

// SelectAll select statement for Orders
func (orderModel *OrderModel) SelectAll() ([]entity.Order, error) {
	rows, err := orderModel.DB.Query("select * from orders")
	if err != nil {
		return nil, err
	}
	var orders []entity.Order
	for rows.Next() {
		var orderID, orderNumber, personID int
		err := rows.Scan(&orderID, &orderNumber, &personID)
		if err != nil {
			return nil, err
		}
		personRes := orderModel.DB.QueryRow("select * from person where id=?", personID)

		var person entity.Person
		err = personRes.Scan(&person.PersonID, &person.Name)

		order := entity.Order{
			OrderID:     orderID,
			OrderNumber: orderNumber,
			Person:      &person,
		}

		if err != nil {
			return nil, err
		}
		// log this in the logging engine

		orders = append(orders, order)
	}
	return orders, nil
}
