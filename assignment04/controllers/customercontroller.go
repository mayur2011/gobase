package controllers

import (
	"encoding/json"
	"fmt"
	"gobase/assignment04/domain"
	"net/http"

	"github.com/gorilla/mux"
)

type CustomerController struct {
	// Dependencies and States
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
		//fmt.Println(err.Error())
		return nil, http.StatusInternalServerError, err
	}
	return customer, http.StatusCreated, nil
}

func (cc CustomerController) GetAllCustomers(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	customers, err := cc.Store.GetAllCustomers()
	if err != nil {
		return nil, http.StatusOK, fmt.Errorf("failed to fetch all customers")
	}
	return customers, http.StatusOK, nil
}

func (cc CustomerController) GetCustomerById(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	vars := mux.Vars(r)
	id := vars["id"]
	customer, err := cc.Store.GetCustomerById(id)
	if err != nil {
		return nil, http.StatusOK, fmt.Errorf("failed to get Customer info")
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
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	return nil, http.StatusAccepted, err
}

func (cc CustomerController) DeleteCustomer(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	vars := mux.Vars(r)
	id := vars["id"]
	err := cc.Store.Delete(id)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	return nil, http.StatusAccepted, err
}
