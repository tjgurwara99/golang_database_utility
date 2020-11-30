package model

import (
	"database/sql"

	"github.com/tjgurwara99/golang_database_utility/entity"
)

// PersonModel Model for Person
type PersonModel struct {
	DB *sql.DB
}

// SelectAll select all statement for person table
func (personModel *PersonModel) SelectAll() ([]entity.Person, error) {
	rows, err := personModel.DB.Query("select * from person")
	if err != nil {
		return nil, err
	}

	var persons []entity.Person
	for rows.Next() {
		var person entity.Person
		err := rows.Scan(&person.PersonID, &person.Name)
		if err != nil {
			return nil, err
		}
		persons = append(persons, person)
	}
	return persons, nil
}
