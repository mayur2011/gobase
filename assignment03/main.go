package main

import (
	"gobase/assignment03/controller"
	"gobase/assignment03/domain"
	"gobase/assignment03/mapstore"
)

func main() {
	cc := controller.CustomerController{
		Store: mapstore.NewMapStore(),
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

/*
-- Need to change lowercase "store" attribute to Camel case otherwise --store will be accessable

*/
