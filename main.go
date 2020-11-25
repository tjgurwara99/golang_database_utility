package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

// Person
type Person struct {
	person_id int
	name      string
}

// Orders
type Orders struct {
	order_id     int
	order_number int
	Person
}

// OpenDatabase returns database object if successful
func OpenDatabase(databaseProgram string, username string, password string, hostname string, databaseName string) (*sql.DB, error) {
	db, err := sql.Open(databaseProgram, username+":"+password+"@tcp("+hostname+")/"+databaseName)
	return db, err
}

func main() {
	db, err := OpenDatabase("mysql", os.Getenv("USER"), os.Getenv("PASSWORD"), "127.0.0.1", os.Getenv("DATABASE_NAME"))

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
		var order Orders
		var person_id int
		err := res.Scan(&order.order_id, &order.order_number, &person_id)

		if err != nil {
			panic(err)
		}

		person_res := db.QueryRow("Select * from person where id = ?", person_id)

		err = person_res.Scan(&order.person_id, &order.name)

		if err != nil {
			panic(err)
		}

		fmt.Printf("%v\n", order)
	}

}
