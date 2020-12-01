package entity

import (
	"fmt"
	"strconv"
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
	BirthDate   time.Time
	IsManager   bool
	IsOwner     bool
	*Company
}

// String returns the default string value
func (user *User) String() string {
	return fmt.Sprintf("%s, %s", user.LastName, user.FirstName)
}

// NewUser Constructor
func NewUser(username, password, firstName, lastName, email string,
	isSuperUser, isStaff, isActive, isManager, isOwner bool,
	company *Company, birthDate time.Time) (*User, error) {
	user := &User{
		UserName:    username,
		FirstName:   firstName,
		LastName:    lastName,
		Email:       email,
		IsSuperUser: isSuperUser,
		IsStaff:     isStaff,
		IsActive:    isActive,
		DateJoined:  time.Now(),
		Company:     company,
		BirthDate:   birthDate,
		IsManager:   isManager,
		IsOwner:     isOwner,
	}
	pass, err := GeneratePassword(password)
	if err != nil {
		return nil, err
	}
	user.Password = pass
	err = user.Validate()
	if err != nil {
		return nil, ErrInvalidDataInput
	}
	return user, nil
}

// ValidatePasswordMinimumLength Validates that password is at least 8 characters long
func (user *User) ValidatePasswordMinimumLength(password string) error {
	const minLength = 8
	if len(user.Password) < minLength {
		return ErrPasswordLength
	}
	return nil

}

// ValidateUsernamePasswordTooSimilar validates whether password is too similar to username
func (user *User) ValidateUsernamePasswordTooSimilar(password string) error {
	// do some sequence matching - not familiar with the difflib library
	// will get back to it when I have the time
	return nil
}

// ValidateNumericPassword Returns error if password is just numeric only
func (user *User) ValidateNumericPassword(password string) error {
	_, err := strconv.Atoi(password)
	if err != nil {
		return err
	}
	return nil
}

// Validate User input validation checks
func (user *User) Validate() error {
	if user.UserName == "" || user.Email == "" || user.FirstName == "" ||
		user.LastName == "" {
		return ErrInvalidDataInput
	}

	return nil
}

// ValidatePassword Validating User Password with Hash
func (user *User) ValidatePassword(password string) error {
	err := user.ValidatePasswordMinimumLength(password)
	if err != nil {
		return err
	}
	err = user.ValidateUsernamePasswordTooSimilar(password)
	if err != nil {
		return err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
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
