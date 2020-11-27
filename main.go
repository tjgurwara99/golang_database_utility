package main

import (
	"fmt"
	"github.com/tjgurwara99/golang_database_utility/models"
	"github.com/tjgurwara99/golang_database_utility/services"
	"os"
)

func main() {
	db, err := services.OpenDatabase("mysql", os.Getenv("DATABASE_USER"), os.Getenv("DATABASE_PASSWORD"), "127.0.0.1", os.Getenv("DATABASE_NAME"))

	if err != nil {
		panic(err)
	}

	defer db.Close()

	var temp models.Order
	res, err := services.Exec(db, temp)

	if err != nil {
		panic(err)
	}

	defer res.Close()

	for res.Next() {
		var order models.Order
		person := models.Person{
			PersonID: 0,
			Name:     "",
		}
		order.Person = &person
		err := res.Scan(&order.OrderID, &order.OrderNumber, &order.PersonID)

		if err != nil {
			panic(err)
		}

		personRes := db.QueryRow("Select * from person where id = ?", order.PersonID)

		err = personRes.Scan(&order.PersonID, &order.Name)

		if err != nil {
			panic(err)
		}

		fmt.Printf("%v\n", &order)
	}

}
