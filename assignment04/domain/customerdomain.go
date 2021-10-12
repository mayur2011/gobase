package domain

import "errors"

// ErrorIDExists is an error value for duplicate customer id
var ErrorIDExists = errors.New("Customer ID exists")

type Customer struct {
	ID, Name, Email string
}

type CustomerStore interface {
	Create(Customer) error
	Update(string, Customer) error
	Delete(string) error
	GetCustomerById(string) (Customer, error)
	GetAllCustomers() ([]Customer, error)
}
