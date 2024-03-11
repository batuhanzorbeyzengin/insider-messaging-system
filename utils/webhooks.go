// utils/webhooks.go
package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github/batuhanzorbeyzengin/insider-messaging-system/database/models"
)

const webhookURL = "https://webhook.site/2baa69ce-0241-4a7b-8fc9-4845a3206e9c"

type WebhookResponse struct {
	Message   string `json:"message"`
	MessageID string `json:"messageId"`
}

// SendMessageToWebhook sends a message to the webhook and returns the response.
func SendMessageToWebhook(message models.Message) (*WebhookResponse, error) {
	payload, err := json.Marshal(map[string]string{
		"to":      message.RecipientPhone,
		"content": message.Content,
	})
	if err != nil {
		return nil, fmt.Errorf("error marshaling message payload: %v", err)
	}

	req, err := http.NewRequest("POST", webhookURL, bytes.NewBuffer(payload))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending message: %v", err)
	}
	defer resp.Body.Close()

	var webhookResponse WebhookResponse
	err = json.NewDecoder(resp.Body).Decode(&webhookResponse)
	if err != nil {
		return nil, fmt.Errorf("error decoding response body: %v", err)
	}

	return &webhookResponse, nil
}
