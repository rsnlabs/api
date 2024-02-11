package textRoutes

import (
	"bytes"
	"encoding/json"
	"os"
  "github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"api/src/middleware"
)

func ClaudeHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	if len(body) == 0 {
		noPromptResponse := map[string]string{"message": "No prompt was provided"}
		jsonResponse, err := json.Marshal(noPromptResponse)
		if err != nil {
			http.Error(w, "Error creating JSON response", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(jsonResponse)
		return
	}

	openchatURL := "https://api.rsnai.org/api/v1/user/claude"
	openchatBearerKey := os.Getenv("APIKEY")

	req, err := http.NewRequest("POST", openchatURL, bytes.NewBuffer(body))
	if err != nil {
		http.Error(w, "Error creating Claude request", http.StatusInternalServerError)
		return
	}

	req.Header.Set("Authorization", "Bearer "+openchatBearerKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "Error making request to Claude API", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	openchatResponse, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Error reading Claude response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	w.Write(openchatResponse)
}

func RegisterClaudeRoute(r *mux.Router) {
	r.Handle("/api/claude", middleware.AuthMiddleware(http.HandlerFunc(ClaudeHandler))).Methods("POST")
}