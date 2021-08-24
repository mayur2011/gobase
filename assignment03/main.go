package main

import (
	"fmt"
	"gobase/assignment03/domain"
	"gobase/assignment03/mapstore"
)

type CustomerController struct {
	store domain.CustomerStore
}

func (cc CustomerController) Add(c domain.Customer) {
	err := cc.store.Create(c)
	if err != nil {
		fmt.Println("Failed to create", err)
	}
	fmt.Println("Customer has been added")
}

func (cc CustomerController) GetAllCustomers() {
	customers, err := cc.store.GetAllCustomers()
	if err != nil {
		fmt.Println("Failed to fetch all customers", err)
	}
	fmt.Println(customers)
}

func main() {
	cc := CustomerController{
		store: mapstore.NewMapStore(),
	}
	customer := domain.Customer{
		ID:    "C101",
		Name:  "Raju",
		Email: "raju-tech@gnn.in",
	}
	cc.Add(customer)
	customer = domain.Customer{
		ID:    "C202",
		Name:  "Maju",
		Email: "maju-tech@gnn.in",
	}
	cc.Add(customer)
	cc.GetAllCustomers()
}
