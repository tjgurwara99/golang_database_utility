package entity

import (
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
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

// ValidatePassword Validating User Password with Hash
func (user *User) ValidatePassword(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return err
	}
	return nil
}

// GeneratePassword Generates the hash string to store in the database
func GeneratePassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 20) // 20: Not sure how intensive I want the hashing to be - check later
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
