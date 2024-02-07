// sms.go
package direct7

import (
	"fmt"
	"log"
)

// SMS struct represents the SMS service in Go.
type SMS struct {
	client *Client
}

// NewSMS creates a new instance of the SMS service.
func NewSMS(client *Client) *SMS {
	return &SMS{client: client}
}

// SendMessages sends one or more messages to a single/multiple recipient(s).
func (s *SMS) SendMessages(messages []Message, originator, reportURL, scheduleTime string) ([]byte, error) {
	payload := MessagePayload{
		Messages: messages,
		MessageGlobals: MessageGlobals{
			Originator:   originator,
			ReportURL:    reportURL,
			ScheduleTime: scheduleTime,
		},
	}
	response, err := s.client.Post("/messages/v1/send", true, payload)
	if err != nil {
		return nil, err
	}
	log.Println("Messages sent successfully.")
	return string(response), nil
}

// GetStatus retrieves the status for a message request.
func (s *SMS) GetStatus(requestID string) ([]byte, error) {
	response, err := s.client.Get(fmt.Sprintf("/report/v1/message-log/%s", requestID), nil)
	if err != nil {
		return nil, err
	}
	log.Println("Message status retrieved successfully.")
	return string(response), nil
}

// MessagePayload represents the payload structure for sending messages.
type MessagePayload struct {
	Messages       []Message       `json:"messages"`
	MessageGlobals MessageGlobals  `json:"message_globals"`
}

// MessageGlobals represents global parameters for a message request.
type MessageGlobals struct {
	Originator   string `json:"originator"`
	ReportURL    string `json:"report_url,omitempty"`
	ScheduleTime string `json:"schedule_time,omitempty"`
}

// Message represents a message in Go.
type Message struct {
    Recipients []string `json:"recipients"`
    Content    string   `json:"content"`
    Unicode    string   `json:"unicode"`
}