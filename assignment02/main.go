package main

import (
	"fmt"
	"gobase/assignment02/domain"
	"gobase/assignment02/mapstore"
)

/* Explicit dependency and declarative programming
that hides dependent logic */
type CustomerController struct {
	store domain.CustomerStore
}

func (cc CustomerController) Add(c domain.Customer) {
	err := cc.store.Create(c)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("New Customer has been added")
}

func (cc CustomerController) GetById(id string) {
	customer, err := cc.store.GetById(id)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Customer info for ID is...\n", customer)
}

func (cc CustomerController) GetAll() {
	customers, err := cc.store.GetAll()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("All existing customers are\n", customers)
}

func main() {
	controller := CustomerController{ //Facade
		store: mapstore.NewMapStore(), //inject the dependency
	}
	customer := domain.Customer{
		ID:        "CUST101",
		Name:      "JPM",
		Email:     "customer01@gmail.com",
		Country:   "India",
		MobileNum: 1113612345,
	}
	controller.Add(customer)
	customer = domain.Customer{
		ID:        "CUST102",
		Name:      "JPM",
		Country:   "India",
		MobileNum: 1213612345,
	}
	controller.Add(customer)
	controller.Add(customer)
	controller.GetAll()
	customer = domain.Customer{
		ID:        "CUST103",
		Name:      "TPM",
		Email:     "",
		Country:   "India",
		MobileNum: 1313612345,
	}
	controller.Add(customer)
	controller.GetAll()
	controller.GetById("CUST102")
	controller.GetById("CUST103")
}

/*
1. Does CustomerController have isExistMethod to check something ?
Ans:

2. Empty string of Email is not being showed by map while displaying
Ans:

*/
