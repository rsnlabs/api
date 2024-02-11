package middleware

import (
	"encoding/json"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"net/http"

	"api/src/utils"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(utils.ColorRed + "Error loading .env file" + utils.ColorReset)
	}
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")

		if token == "" || !strings.HasPrefix(token, "Bearer ") {
			unauthorizedResponse := map[string]string{"error": "Unauthorized"}
			jsonResponse, err := json.Marshal(unauthorizedResponse)
			if err != nil {
				http.Error(w, utils.ColorRed+"Error creating JSON response"+utils.ColorReset, http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write(jsonResponse)
			return
		}

		actualToken := strings.TrimPrefix(token, "Bearer ")

		if !isValidToken(actualToken) {
			unauthorizedResponse := map[string]string{"error": "Unauthorized"}
			jsonResponse, err := json.Marshal(unauthorizedResponse)
			if err != nil {
				http.Error(w, utils.ColorRed+"Error creating JSON response"+utils.ColorReset, http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write(jsonResponse)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func isValidToken(token string) bool {
	authKey := os.Getenv("AUTH_KEY")
	log.Println(utils.ColorGreen + "Loaded AUTH_KEY: " + authKey + utils.ColorReset)
	return token == authKey
}