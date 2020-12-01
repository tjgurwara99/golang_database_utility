package main

import (
	"fmt"

	"github.com/tjgurwara99/golang_database_utility/config"
	"github.com/tjgurwara99/golang_database_utility/model"
	"github.com/tjgurwara99/golang_database_utility/service"
)

func main() {
	db, err := service.OpenDatabase(config.GetConfs())

	if err != nil {
		panic(err)
	}

	defer db.Close()

	orderModel := model.OrderModel{DB: db}

	orders, err := orderModel.SelectAll()
	if err != nil {
		fmt.Println(err)
	}

	for _, value := range orders {
		fmt.Println(&value)
	}

}
