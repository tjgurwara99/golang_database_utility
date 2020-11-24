package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

// OpenDatabase returns database object if successful
func OpenDatabase(databaseProgram string, username string, password string, hostname string, databaseName string) (*sql.DB, error) {
	db, err := sql.Open(databaseProgram, username+":"+password+"@tcp("+hostname+")/"+databaseName)
	return db, err
}

func main() {
	db, err := OpenDatabase("mysql", os.Getenv("USERNAME"), os.Getenv("PASSWORD"), "127.0.0.1", os.Getenv("DATABASE_NAME"))
	if err != nil {
		panic(err)
	}

	defer db.Close()

}
