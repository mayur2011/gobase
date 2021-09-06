package main

import (
	"fmt"
	"gobase/assignment03/router"
	"net/http"
	"os"
)

func main() {
	router := router.InitRoutes()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	fmt.Println("Launching the app, visit localhost:8000/")
	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		fmt.Print(err)
	}
}
