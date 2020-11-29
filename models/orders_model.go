package models

import (
	"database/sql"
	"github.com/tjgurwara99/golang_database_utility/entities"
)

//OrderModel Order Entity Model
type OrderModel struct {
	DB *sql.DB
}

// SelectAll select statement for Orders
func (orderModel OrderModel) SelectAll() ([]entities.Order, error) {
	rows, err := orderModel.DB.Query("select * from orders")
	if err != nil {
		return nil, err
	}
	var orders []entities.Order
	for rows.Next() {
		var orderID, orderNumber, personID int
		err := rows.Scan(&orderID, &orderNumber, &personID)
		if err != nil {
			return nil, err
		}
		personRes := orderModel.DB.QueryRow("select * from person where id=?", personID)

		var personName string
		err = personRes.Scan(&personID, &personName)

		order := entities.Order{
			OrderID:     orderID,
			OrderNumber: orderNumber,
			Person:      &entities.Person{personID, personName},
		}

		if err != nil {
			return nil, err
		}
		// log this in the logging engine

		orders = append(orders, order)
	}
	return orders, nil
}
