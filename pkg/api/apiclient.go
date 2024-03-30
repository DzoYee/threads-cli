package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"threads/pkg/models"
)

// SendTitle sends the title to the API
func SendTitle(title string) error {
	apiURL := "http://yourapi.com/endpoint" // Change to your API's URL
	requestPayload := models.TitleRequest{Title: title}
	payloadBytes, err := json.Marshal(requestPayload)
	if err != nil {
		return err
	}

	resp, err := http.Post(apiURL, "application/json", bytes.NewBuffer(payloadBytes))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("received non-ok response code: %d", resp.StatusCode)
	}
	return nil
}
