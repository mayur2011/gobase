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
	return router
}

// InitRoutes registers all routes for the application.
func InitRoutes() *mux.Router {
	router := mux.NewRouter()
	router = SetCustomerRoutes(router)
	return router
}
