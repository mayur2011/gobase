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
	fmt.Println("New Customer has been created")
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
		ID:   "CUST101",
		Name: "JPM",
	}
	controller.Add(customer)
	customer = domain.Customer{
		ID:   "CUST102",
		Name: "JPM",
	}
	controller.Add(customer)
	controller.GetAll()
	customer = domain.Customer{
		ID:   "CUST103",
		Name: "TPM",
	}
	controller.Add(customer)
	customer = domain.Customer{
		ID:   "CUST104",
		Name: "TPM",
	}
	controller.Add(customer)
	controller.GetAll()
}
