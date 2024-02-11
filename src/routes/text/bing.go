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

func BingHandler(w http.ResponseWriter, r *http.Request) {
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

	bingURL := "https://api.rsnai.org/api/v1/user/bing"
	bingBearerKey := os.Getenv("APIKEY")

	req, err := http.NewRequest("POST", bingURL, bytes.NewBuffer(body))
	if err != nil {
		http.Error(w, "Error creating GPT request", http.StatusInternalServerError)
		return
	}

	req.Header.Set("Authorization", "Bearer "+bingBearerKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "Error making request to Bing API", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	bingResponse, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Error reading Bing response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	w.Write(bingResponse)
}

func RegisterBingRoute(r *mux.Router) {
	r.Handle("/api/bing", middleware.AuthMiddleware(http.HandlerFunc(BingHandler))).Methods("POST")
}