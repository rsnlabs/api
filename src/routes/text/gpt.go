package textRoutes

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"api/src/middleware"
)

func GptHandler(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error reading request body"})
		return
	}

	if len(body) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "No prompt was provided"})
		return
	}

	APIURL := "https://api.rnilaweera.lk/api/v1/user/gpt"
	BearerKey := os.Getenv("APIKEY")

	req, err := http.NewRequest("POST", APIURL, bytes.NewBuffer(body))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating GPT request"})
		return
	}

	req.Header.Set("Authorization", "Bearer "+BearerKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error making request to GPT API"})
		return
	}
	defer resp.Body.Close()

	Response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error reading GPT response"})
		return
	}

	c.Data(resp.StatusCode, "application/json", Response)
}

func RegisterGptRoute(r *gin.Engine) {
	r.POST("/api/gpt", middleware.AuthMiddleware(), GptHandler)
}