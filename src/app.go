package main

import (
	"log"
	"net/http"
	"os"
	"github.com/gorilla/mux"
	"api/src/routes"
)

func main() {
	r := mux.NewRouter()

	routes.RegisterHomeRoute(r)
	
	routes.RegisterGPTRoute(r)
	routes.RegisterOpenChatRoute(r)
	routes.RegisterBardRoute(r)
	routes.RegisterLlaMaRoute(r)
	routes.RegisterMixtralRoute(r)

	logger := log.New(os.Stdout, "Server: ", log.LstdFlags)

	logger.Printf("Server is listening on :5000")

	if err := http.ListenAndServe(":5000", r); err != nil {
		logger.Fatal("Error starting server: ", err)
	}
}