package mapstore

import (
	"fmt"
	"gobase/assignment02/domain"
	"log"
)

//struct mapstore to implement CustomerStore interface
type MapStore struct {
	store map[string]domain.Customer //An in-memory store in a map
}

//factory methods gives a new instance of MapStore
//this is for caller packages, not for mapstore
func NewMapStore() *MapStore {
	return &MapStore{store: make(map[string]domain.Customer)}
}

//implementing interface methods of domain.CustomerStore
func (ms *MapStore) isCustomerExists(customer domain.Customer) bool {
	_, ok := ms.store[customer.ID]
	return ok
}

func (ms *MapStore) Create(customer domain.Customer) error {
	if ms.isCustomerExists(customer) {
		return fmt.Errorf("customer already exists")
	}
	ms.store[customer.ID] = customer
	log.Println("customer is added")
	return nil
}

func (ms *MapStore) GetAll() ([]domain.Customer, error) {
	var c []domain.Customer
	for k, v := range ms.store {
		fmt.Println(k, "", v)
		c = append(c, ms.store[k])
	}
	log.Println("???")
	return c, nil
}
