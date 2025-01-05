package main

import (
	"fmt"
	"net/http"
	"reservation-system/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	routes.RegisterRoutes(r)
	if err := http.ListenAndServe(":8080", r); err != nil {
		fmt.Println("Error: ", err)
	}
}
