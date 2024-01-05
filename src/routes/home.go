package routes

import (
	"net/http"
	"github.com/gorilla/mux"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the API!"))
}

func RegisterHomeRoute(r *mux.Router) {
	r.HandleFunc("/", HomeHandler).Methods("GET")
}