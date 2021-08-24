package controller

import (
	"fmt"
	"gobase/assignment03/domain"
)

type CustomerController struct {
	Store domain.CustomerStore
}

func (cc CustomerController) Add(c domain.Customer) {
	err := cc.Store.Create(c)
	if err != nil {
		fmt.Println("Failed to create", err)
	}
	fmt.Println("Customer has been added")
}

func (cc CustomerController) GetAllCustomers() {
	customers, err := cc.Store.GetAllCustomers()
	if err != nil {
		fmt.Println("Failed to fetch all customers", err)
	}
	fmt.Println(customers)
}
