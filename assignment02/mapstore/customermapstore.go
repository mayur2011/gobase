package mapstore

import (
	"fmt"
	"gobase/assignment02/domain"
)

//struct mapstore to implement CustomerStore interface
type MapStore struct {
	store map[string]domain.Customer //An in-memory store in a map
}

//factory methods gives a new instance of MapStore, this is for caller packages, not for mapstore
func NewMapStore() *MapStore {
	return &MapStore{store: make(map[string]domain.Customer)}
}

//implementing interface methods of domain.CustomerStore
func (ms *MapStore) isCustomerExists(id string) bool {
	_, ok := ms.store[id]
	return ok
}

func (ms *MapStore) Create(customer domain.Customer) error {
	if ms.isCustomerExists(customer.ID) {
		return fmt.Errorf("customer already exists in mapstore")
	}
	ms.store[customer.ID] = customer
	return nil
}

func (ms *MapStore) Update(id string, customer domain.Customer) error {
	if ms.isCustomerExists(id) {
		ms.store[id] = customer
		return nil
	}
	return fmt.Errorf("update failed, customer id does not exist")
}

func (ms *MapStore) Delete(id string) error {
	if ms.isCustomerExists(id) {
		delete(ms.store, id)
		return nil
	}
	return fmt.Errorf("delete failed, customer id does not exist")
}

func (ms *MapStore) GetById(id string) (domain.Customer, error) {
	if ms.isCustomerExists(id) {
		c := ms.store[id]
		return c, nil
	}
	return domain.Customer{}, fmt.Errorf("customer does not exist for this id")
}

func (ms *MapStore) GetAll() ([]domain.Customer, error) {
	var c []domain.Customer
	for k := range ms.store {
		c = append(c, ms.store[k])
	}
	return c, nil
}

/*
Questions:
1. is it okay for GetById func to return empty customer interface (line no. 53) ?
2.
*/
