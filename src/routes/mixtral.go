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

func MixtralHandler(w http.ResponseWriter, r *http.Request) {
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

	mixtralURL := "https://ai.rnilaweera.ovh/api/v1/user/mixtral"
	mixtralBearerKey := os.Getenv("APIKEY")

	req, err := http.NewRequest("POST", mixtralURL, bytes.NewBuffer(body))
	if err != nil {
		http.Error(w, "Error creating Mixtral request", http.StatusInternalServerError)
		return
	}

	req.Header.Set("Authorization", "Bearer "+mixtralBearerKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "Error making request to Mixtral API", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	mixtralResponse, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Error reading Mixtral response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	w.Write(mixtralResponse)
}

func RegisterMixtralRoute(r *mux.Router) {
	r.Handle("/mixtral", middleware.AuthMiddleware(http.HandlerFunc(MixtralHandler))).Methods("POST")
}