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
	IsSuperuser bool
	Username    string
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
	IsSuperuser, isStaff, isActive, isManager, isOwner bool,
	company *Company, birthDate time.Time) (*User, error) {
	user := &User{
		Username:    username,
		FirstName:   firstName,
		LastName:    lastName,
		Email:       email,
		IsSuperuser: IsSuperuser,
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
func ValidatePasswordMinimumLength(password string) error {
	const minLength = 8
	if len(password) < minLength {
		return ErrPasswordLength
	}
	return nil

}

// ValidateUsernamePasswordTooSimilar validates whether password is too similar to Username
func ValidateUsernamePasswordTooSimilar(password string) error {
	// do some sequence matching - not familiar with the difflib library
	// will get back to it when I have the time
	// TODO: Read the documentation of difflib and figure out a way to
	// do sequence matching with different fields of the user.
	// I may need to add another pointer argument to the User.
	return nil
}

// ValidateNumericPassword Returns error if password is just numeric only
func ValidateNumericPassword(password string) error {
	//match, err := regexp.MatchString(`^[0-9 ]+$`, password)
	//if err != nil {
	//	return err
	//}
	//if match {
	//	return ErrNumericPassword
	//}

	return nil
}

// Validate User input validation checks
func (user *User) Validate() error {
	if user.Username == "" || user.Email == "" || user.FirstName == "" ||
		user.LastName == "" {
		return ErrInvalidDataInput
	}

	return nil
}

// CheckPassword Validating This is to log the User in with Password
func (user *User) CheckPassword(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return err
	}
	return nil
}

// GeneratePassword Generates the hash string to store in the database
func GeneratePassword(password string) (string, error) {
	// Don't proceed forward if the password is numeric
	err := ValidateNumericPassword(password)
	if err != nil {
		return "", err
	}
	// Don't proceed forward if the password is too
	// similar to the username - NOT WORKING RIGHT NOW
	// Look into the library difflib (similar to
	// python's difflib I think but read the
	// documentation to be sure.)
	err = ValidateUsernamePasswordTooSimilar(password)
	if err != nil {
		return "", err
	}
	// Hash the password and return the hashed password
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10) // 20: Not sure how intensive I want the hashing to be - check later
	if err != nil {
		fmt.Println("It reached here")
		return "", err
	}

	// TODO: Read NIST standards for passwords
	return string(hash), nil
}
