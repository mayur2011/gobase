package main

import (
	"fmt"
	"gobase/assignment02/domain"
	"gobase/assignment02/mapstore"
)

type CustomerController struct {
	store domain.CustomerStore
}

func (cc CustomerController) Add(c domain.Customer) {
	err := cc.store.Create(c)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("New Customer has been created")
}

func main() {
	controller := CustomerController{
		store: mapstore.NewMapStore(),
	}
	customer := domain.Customer{
		ID:   "CUST101",
		Name: "JPM",
	}
	controller.Add(customer)

}
