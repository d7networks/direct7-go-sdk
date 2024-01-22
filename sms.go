
// sms.go
package direct7

import (
	"fmt"
	"log"
)

type SMS struct {
	client *Client
}

func NewSMS(client *Client) *SMS {
	return &SMS{
		client: client,
	}
}

func (s *SMS) SendMessage(recipients []string, content, originator, reportURL string, unicode bool) (map[string]interface{}, error) {
	message := map[string]interface{}{
		"channel":      "sms",
		"recipients":   recipients,
		"content":      content,
		"msg_type":     "text",
		"data_coding":  "unicode",
	}
	if !unicode {
		message["data_coding"] = "text"
	}

	messageGlobals := map[string]interface{}{
		"originator": originator,
		"report_url": reportURL,
	}

	params := map[string]interface{}{
		"messages":         []map[string]interface{}{message},
		"message_globals":  messageGlobals,
	}

	response, err := s.client.post(s.client.host(), "/messages/v1/send", true, params)
	if err != nil {
		log.Printf("Error sending SMS: %v\n", err)
		return nil, err
	}

	log.Println("Message sent successfully.")
	return response, nil
}

func (s *SMS) GetStatus(requestID string) (map[string]interface{}, error) {
	response, err := s.client.get(s.client.host(), fmt.Sprintf("/report/v1/message-log/%s", requestID), nil)
	if err != nil {
		log.Printf("Error getting SMS status: %v\n", err)
		return nil, err
	}

	log.Println("Message status retrieved successfully.")
	return response, nil
}
