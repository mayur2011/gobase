package controller

import (
	"encoding/json"
	"fmt"
	"gobase/assignment03/domain"
	"net/http"

	"github.com/gorilla/mux"
)

type CustomerController struct {
	Store domain.CustomerStore
}

func (cc CustomerController) PostCustomer(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	var customer domain.Customer
	// decode the incoming request
	err := json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		return nil, http.StatusBadRequest, fmt.Errorf("unable to decode JSON request body")
	}
	err = cc.Store.Create(customer)
	if err != nil {
		fmt.Println("Failed to create the customer", err)
	}
	return customer, http.StatusCreated, nil
}

func (cc CustomerController) GetAllCustomers(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	customers, err := cc.Store.GetAllCustomers()
	if err != nil {
		fmt.Println("Failed to fetch all customers", err)
	}
	return customers, http.StatusOK, nil
}

func (cc CustomerController) GetCustomerById(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	vars := mux.Vars(r)
	id := vars["id"]
	customer, err := cc.Store.GetCustomerById(id)
	if err != nil {
		fmt.Println("Failed to get Customer info", err)
	}
	return customer, http.StatusOK, nil
}

func (cc CustomerController) UpdateCustomer(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	var customer domain.Customer
	vars := mux.Vars(r)
	id := vars["id"]
	err := json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		return nil, http.StatusBadRequest, fmt.Errorf("unable to decode JSON request body")
	}
	err = cc.Store.Update(id, customer)
	return customer, http.StatusAccepted, err
}
