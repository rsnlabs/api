package routes

import (
	"bytes"
	"encoding/json"
	"os"
  "github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"api/src/middleware"
)

func GPTHandler(w http.ResponseWriter, r *http.Request) {
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

	gptURL := "https://ai.rnilaweera.ovh/api/v1/user/gpt"
	gptBearerKey := os.Getenv("APIKEY")

	req, err := http.NewRequest("POST", gptURL, bytes.NewBuffer(body))
	if err != nil {
		http.Error(w, "Error creating GPT request", http.StatusInternalServerError)
		return
	}

	req.Header.Set("Authorization", "Bearer "+gptBearerKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "Error making request to GPT API", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	gptResponse, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Error reading GPT response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	w.Write(gptResponse)
}

func RegisterGPTRoute(r *mux.Router) {
	r.Handle("/gpt", middleware.AuthMiddleware(http.HandlerFunc(GPTHandler))).Methods("POST")
}