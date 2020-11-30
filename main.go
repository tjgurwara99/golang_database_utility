package main

import (
	"fmt"

	"github.com/tjgurwara99/golang_database_utility/config"
	"github.com/tjgurwara99/golang_database_utility/model"
	"github.com/tjgurwara99/golang_database_utility/services"
)

func main() {
	db, err := services.OpenDatabase(config.GetConfs())

	if err != nil {
		panic(err)
	}

	defer db.Close()

	orderModel := model.OrderModel{db}

	orders, err := orderModel.SelectAll()
	if err != nil {
		fmt.Println(err)
	}

	for _, value := range orders {
		fmt.Println(&value)
	}

}
