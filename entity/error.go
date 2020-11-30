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
