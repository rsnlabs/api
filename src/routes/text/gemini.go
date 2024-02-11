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

func GeminiHandler(w http.ResponseWriter, r *http.Request) {
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

	geminiURL := "https://api.rsnai.org/api/v1/user/gemini"
	geminiBearerKey := os.Getenv("APIKEY")

	req, err := http.NewRequest("POST", geminiURL, bytes.NewBuffer(body))
	if err != nil {
		http.Error(w, "Error creating Gemini request", http.StatusInternalServerError)
		return
	}

	req.Header.Set("Authorization", "Bearer "+geminiBearerKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "Error making request to Gemini API", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	geminiResponse, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Error reading Gemini response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	w.Write(geminiResponse)
}

func RegisterGeminiRoute(r *mux.Router) {
	r.Handle("/api/gemini", middleware.AuthMiddleware(http.HandlerFunc(GeminiHandler))).Methods("POST")
}