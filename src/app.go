package main

import (
	"log"
	"github.com/gin-gonic/gin"

	"api/src/routes"
	"api/src/routes/text"
)

func main() {
	router := gin.Default()

	routes.RegisterHomeRoute(router)

	textRoutes.RegisterGptRoute(router)
	textRoutes.RegisterBardRoute(router)
	textRoutes.RegisterGeminiRoute(router)
	textRoutes.RegisterLlaMaRoute(router)
	textRoutes.RegisterCodeLlamaRoute(router)
	textRoutes.RegisterMixtralRoute(router)
	textRoutes.RegisterClaudeRoute(router)

	log.Println("Server is listening on :5000")
	if err := router.Run(":5000"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}