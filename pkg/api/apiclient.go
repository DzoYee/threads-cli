package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"threads/pkg/models"
)

// SendIdea sends the title to the API
func SendTitle(name string) error {
	apiURL := "http://localhost:8080/ideas" // Change to your API's URL
	requestPayload := models.IdeaRequest{Name: name}
	payloadBytes, err := json.Marshal(requestPayload)
	if err != nil {
		return err
	}

	resp, err := http.Post(apiURL, "application/json", bytes.NewBuffer(payloadBytes))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check if the status code is within the range of success codes (200-299)
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return fmt.Errorf("received non-success response code: %d", resp.StatusCode)
	}
	return nil
}
