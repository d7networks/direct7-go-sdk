package direct7

import (
	"fmt"
	"log"
)

// Viber represents the Viber service in Go.
type Viber struct {
	client *Client
}

// NewViber creates a new instance of the Viber service.
func NewViber(client *Client) *Viber {
	return &Viber{client: client}
}

// SendViberMessage sends a Viber message to one or more recipients.
func (v *Viber) SendViberMessage(recipients []string, content, label, originator, callBackURL string) (string, error) {
	message := map[string]interface{}{
		"channel":     "viber",
		"recipients":  recipients,
		"content":     content,
		"label":       label,
	}
	messageGlobals := map[string]interface{}{
		"originator":    originator,
		"call_back_url": callBackURL,
	}

	response, err := v.client.Post("/viber/v1/send", true, map[string]interface{}{"messages": []interface{}{message}, "message_globals": messageGlobals})
	if err != nil {
		return "", err
	}
	log.Println("Message sent successfully.")
	return string(response), nil
}

// GetStatus retrieves the status for a Viber message request.
func (v *Viber) GetStatus(requestID string) (string, error) {
	response, err := v.client.Get(fmt.Sprintf("/report/v1/viber-log/%s", requestID), nil)
	if err != nil {
		return "", err
	}
	log.Println("Message status retrieved successfully.")
	return string(response), nil
}
