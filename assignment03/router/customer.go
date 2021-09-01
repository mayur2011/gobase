package router

import (
	"gobase/assignment03/controller"
	"gobase/assignment03/mapstore"

	"github.com/gorilla/mux"
)

func SetCustomerRoutes(router *mux.Router) *mux.Router {
	customerStore := mapstore.NewMapStore()
	customerController := controller.CustomerController{Store: customerStore}
	router.Handle("/customers", controller.ResponseHandler(customerController.PostCustomer)).Methods("POST")
	router.Handle("/customers", controller.ResponseHandler(customerController.GetAllCustomers)).Methods("GET")
	router.Handle("/customers/{id}", controller.ResponseHandler(customerController.GetCustomerById)).Methods("GET")
	router.Handle("/customer/{id}", controller.ResponseHandler(customerController.UpdateCustomer)).Methods("PUT")
	return router
}

// InitRoutes registers all customer routes for the application.
func InitRoutes() *mux.Router {
	router := mux.NewRouter()
	router = SetCustomerRoutes(router)
	return router
}
