package Routes

import (
	"Roteador/handlers"

	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {

	r := mux.NewRouter()

	api := r.PathPrefix("/api").Subrouter()

	api.HandleFunc("/users", handlers.GetUsers).Methods("GET")
	api.HandleFunc("/users/{id}", handlers.GetUser).Methods("GET")
	api.HandleFunc("/users", handlers.CreateUser).Methods("POST")

	return r
}
