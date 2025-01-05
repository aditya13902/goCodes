package routes

import (
	"net/http"
	"reservation-system/controller"

	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router) {

	router.HandleFunc("/add/Store", controller.AddStore).Methods(http.MethodPost)
	router.HandleFunc("/add/User", controller.AddUser).Methods(http.MethodPost)
	router.HandleFunc("/get", controller.GetUser)
}
