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
func (ms *MapStore) isCustomerExists(id string) bool {
	_, ok := ms.store[id]
	return ok
}

func (ms *MapStore) Create(customer domain.Customer) error {
	cust_id := customer.ID
	if ms.isCustomerExists(cust_id) {
		return fmt.Errorf("customer already exists")
	}
	ms.store[customer.ID] = customer
	log.Println("customer is created")
	return nil
}

func (ms *MapStore) Update(id string, customer domain.Customer) error {
	if ms.isCustomerExists(id) {
		ms.store[id] = customer
		log.Println("customer is updated")
		return nil
	}
	return fmt.Errorf("update failed, customer does not exist for this id")
}

func (ms *MapStore) GetById(id string) (domain.Customer, error) {
	if ms.isCustomerExists(id) {
		c := ms.store[id]
		return c, nil
	}
	return domain.Customer{}, fmt.Errorf("customer does not exist for this id")
}

func (ms *MapStore) GetAll() ([]domain.Customer, error) {
	log.Println("return all customer records from mapstore")
	var c []domain.Customer
	for k := range ms.store {
		//fmt.Println(k, "", v)
		c = append(c, ms.store[k])
	}
	return c, nil
}

/*
Questions:
1. isCustomerExists function should accept a key or customer object ?
2. is it okay when GetById func (else part) is return empty customer interface ?
3.
*/
