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

func (ms *MapStore) GetCustomerById(id string) (domain.Customer, error) {
	if ms.isCustomerExists(id) {
		customer := ms.store[id]
		return customer, nil
	}
	return domain.Customer{}, fmt.Errorf("customer does not exist for this id")
}

func (ms *MapStore) Update(id string, customer domain.Customer) error {
	if ms.isCustomerExists(id) {
		ms.store[id] = customer
		fmt.Println("Customer has been updated")
		return nil
	}
	return fmt.Errorf("customer does not exist for this id")
}

func (ms *MapStore) Delete(id string) error {
	if ms.isCustomerExists(id) {
		delete(ms.store, id)
		fmt.Println("Customer has been deleted")
		return nil
	}
	return fmt.Errorf("customer does not exist for this id")
}
