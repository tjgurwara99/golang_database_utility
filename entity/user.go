package entity

import (
	"fmt"
	"time"
)

// User Entity
type User struct {
	UserID      int64
	Password    string
	FirstName   string
	LastName    string
	LastLogin   time.Time
	IsSuperUser bool
	UserName    string
	Email       string
	IsStaff     bool
	IsActive    bool
	DateJoined  time.Time
	*Company
	BirthDate time.Time
}

// String returns the default string value
func (user *User) String() string {
	return fmt.Sprintf("%s, %s", user.LastName, user.FirstName)
}

// This is the way I set up User Entity. Not the only way to do this but I
// find this separates the business logic and requirement logic so they
// don't mix. Also, the main User database model would be one not two
// (i.e. no AbstractUser model is necessary). We can create it but I find
// it will be a hassle if we're working on a small/medium sized project.
// (This is mostly insprired by how Django works but with my subtle tweaks)
