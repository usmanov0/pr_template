package api

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Router(controller *UserController) http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/users", controller.CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", controller.GetUserById).Methods("GET")
	router.HandleFunc("/users", controller.GetUserById).Methods("GET")
	router.HandleFunc("/users", controller.PutUser).Methods("PUT")
	router.HandleFunc("/users/{id}", controller.DeleteUser).Methods("DELETE")

	log.Println("Server started on: 8080")
	return router
}
