package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	"api/src/routes"
	"api/src/routes/text"
	"api/src/utils"
)

func main() {
	r := mux.NewRouter()

	routes.RegisterHomeRoute(r)

	textRoutes.RegisterGPTRoute(r)
	textRoutes.RegisterOpenChatRoute(r)
	textRoutes.RegisterBardRoute(r)
	textRoutes.RegisterGeminiRoute(r)
	textRoutes.RegisterBingRoute(r)
	textRoutes.RegisterLlaMaRoute(r)
	textRoutes.RegisterCodeLlamaRoute(r)
	textRoutes.RegisterMixtralRoute(r)
	textRoutes.RegisterClaudeRoute(r)

	logger := log.New(os.Stdout, utils.ColorGreen+"Server: "+utils.ColorReset, log.LstdFlags)

	logger.Printf(utils.ColorGreen + "Server is listening on :5000" + utils.ColorReset)

	if err := http.ListenAndServe(":5000", r); err != nil {
		logger.Fatal(utils.ColorRed + "Error starting server: " + err.Error() + utils.ColorReset)
	}
}