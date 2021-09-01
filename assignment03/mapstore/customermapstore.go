package mapstore

import (
	"fmt"
	"gobase/assignment03/domain"
)

type MapStore struct {
	store map[string]domain.Customer
}

//Factory method gives a new instance of MapStore, Also this is for caller packages, not for mapstoer itself
func NewMapStore() *MapStore {
	return &MapStore{store: make(map[string]domain.Customer)}
}

func (ms *MapStore) isCustomerExists(id string) bool {
	_, ok := ms.store[id]
	return ok
}

func (ms *MapStore) Create(customer domain.Customer) error {
	if ms.isCustomerExists(customer.ID) {
		return fmt.Errorf("customer exists")
	}
	ms.store[customer.ID] = customer
	fmt.Println("Customer has been created")
	return nil
}

func (ms *MapStore) GetAllCustomers() ([]domain.Customer, error) {
	var customers []domain.Customer
	for k := range ms.store {
		customers = append(customers, ms.store[k])
	}
	return customers, nil
}
