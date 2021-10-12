package router

import (
	"gobase/assignment04/controllers"
	"gobase/assignment04/mapstore"

	"github.com/gorilla/mux"
)

func SetCustomerRoutes(router *mux.Router) *mux.Router {
	customerStore := mapstore.NewMapStore()
	customerController := controllers.CustomerController{Store: customerStore}
	router.Handle("/customers", controllers.ResponseHandler(customerController.PostCustomer)).Methods("POST")
	router.Handle("/customers", controllers.ResponseHandler(customerController.GetAllCustomers)).Methods("GET")
	router.Handle("/customers/{id}", controllers.ResponseHandler(customerController.GetCustomerById)).Methods("GET")
	router.Handle("/customer/{id}", controllers.ResponseHandler(customerController.UpdateCustomer)).Methods("PUT")
	router.Handle("/customer/{id}", controllers.ResponseHandler(customerController.DeleteCustomer)).Methods("DELETE")
	return router
}

// InitRoutes registers all customer routes for the application.
func InitRoutes() *mux.Router {
	router := mux.NewRouter()
	router = SetCustomerRoutes(router)
	return router
}
