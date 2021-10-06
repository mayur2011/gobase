package main

import (
	"fmt"
	"gobase/testing/httpbdd/controllers"
	"gobase/testing/httpbdd/store"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func setUserRoutes() *mux.Router {
	r := mux.NewRouter()
	// Resolve dependencies
	userStore := store.NewUserMapStore()
	controller := controllers.UserController{
		Store: userStore, //Injecting dependencies
	}
	r.HandleFunc("/users", controller.CreateUser).Methods("POST")
	r.HandleFunc("/users", controller.GetUsers).Methods("GET")
	return r
}

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8011"
	}
	fmt.Println("Launching the app, visit localhost:8011/")
	err := http.ListenAndServe(":"+port, setUserRoutes())
	if err != nil {
		fmt.Print(err)
	}
}
