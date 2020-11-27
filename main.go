package main

import (
	"fmt"
	"github.com/tjgurwara99/golang_database_utility/models"
	"github.com/tjgurwara99/golang_database_utility/services"
	"os"
)

func main() {
	db, err := services.OpenDatabase("mysql", os.Getenv("USER"), os.Getenv("PASSWORD"), "127.0.0.1", os.Getenv("DATABASE_NAME"))

	if err != nil {
		panic(err)
	}

	defer db.Close()

	res, err := db.Query("SELECT * FROM orders")

	if err != nil {
		panic(err)
	}

	defer res.Close()

	for res.Next() {
		var order models.Order
		var personID int
		err := res.Scan(&order.OrderID, &order.OrderNumber, &personID)

		if err != nil {
			panic(err)
		}

		personRes := db.QueryRow("Select * from person where id = ?", personID)

		err = personRes.Scan(&order.PersonID, &order.Name)

		if err != nil {
			panic(err)
		}

		fmt.Printf("%v\n", order)
	}

}
