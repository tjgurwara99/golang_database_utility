package entity

import "fmt"

// Person Model
type Person struct {
	PersonID int
	Name     string
}

func (person *Person) String() string {
	return fmt.Sprintf("{PersonID: %d, Name: %s", person.PersonID, person.Name)
}
