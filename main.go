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

	orderModel := models.OrderModel{db}

	orders, err := orderModel.SelectAll()
	if err != nil {
		fmt.Println(err)
	}

	for _, value := range orders {
		fmt.Println(value)
	}

}
