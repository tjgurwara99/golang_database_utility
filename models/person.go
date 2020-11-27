package models

import "fmt"

// Person Model
type Person struct {
	PersonID int
	Name     string
	Orders   []Order
}

func (person *Person) String() string {
	return fmt.Sprintf("{PersonID: %d, Name: %s, Orders: }", person.PersonID, person.Name)
}
