package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{"message": "Lista de usuários"})
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	json.NewEncoder(w).Encode(map[string]string{
		"id": params["id"],
	})
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Usuário criado",
	})
}
