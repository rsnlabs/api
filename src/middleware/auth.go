package middleware

import (
	"encoding/json"
	"log"
	"github.com/joho/godotenv"
	"net/http"
	"os"
	"strings"
)

func init() {
    if err := godotenv.Load(); err != nil {
        log.Fatal("Error loading .env file")
    }
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")

		if token == "" || !strings.HasPrefix(token, "Bearer ") {
			unauthorizedResponse := map[string]string{"error": "Unauthorized"}
			jsonResponse, err := json.Marshal(unauthorizedResponse)
			if err != nil {
				http.Error(w, "Error creating JSON response", http.StatusInternalServerError)
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
				http.Error(w, "Error creating JSON response", http.StatusInternalServerError)
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
    log.Println("Loaded AUTH_KEY:", authKey)
    return token == authKey
}