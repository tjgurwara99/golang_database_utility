package entity

import "errors"

// ErrUserNotFound User is not found error
var ErrUserNotFound = errors.New("User Not Found")

// ErrUserNotActive User is not active
var ErrUserNotActive = errors.New("User is inactive")

// ErrUsernamePassword Username or Password is incorrect
var ErrUsernamePassword = errors.New("Username or Password is incorrect")

// ErrSomethingWrong A generic error - not good to do
// this but will change them to reflect our usecase
var ErrSomethingWrong = errors.New("Something went wrong")

// ErrInvalidDataInput Data in the form is invalid
var ErrInvalidDataInput = errors.New("Invalid inputs")

// ErrPasswordLength Password Too Small
var ErrPasswordLength = errors.New("Password too small")

// ErrNumericPassword Password contains only numbers
var ErrNumericPassword = errors.New("Password only contains numbers")
