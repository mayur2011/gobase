package domain

import "errors"

var ErrorEmailsExists = errors.New("email id exists") //error strings should not be capitalized

type User struct {
	FirstName string
	LastName  string
	Email     string
}

type UserStore interface {
	GetUsers() []User
	AddUser(User) error
}
