package direct7

import (
	"fmt"
	"log"
)

type SMS struct {
	client *Client
}

func NewSMS(client *Client) *SMS {
	return &SMS{client: client}
}

func (s *SMS) SendMessages(messages []Message, originator, reportURL, scheduleTime, tag string) (string, error) {
	payload := MessagePayload{
		Messages: messages,
		MessageGlobals: MessageGlobals{
			Originator:   originator,
			ReportURL:    reportURL,
			ScheduleTime: scheduleTime,
			Tag: tag,
		},
	}
	response, err := s.client.Post("/messages/v1/send", true, payload)
	if err != nil {
		return "", err
	}
	log.Println("Messages sent successfully.")
	return string(response), nil
}

func (s *SMS) GetStatus(requestID string) (string, error) {
	response, err := s.client.Get(fmt.Sprintf("/report/v1/message-log/%s", requestID), nil)
	if err != nil {
		return "", err
	}
	log.Println("Message status retrieved successfully.")
	return string(response), nil
}

type MessagePayload struct {
	Messages       []Message       `json:"messages"`
	MessageGlobals MessageGlobals  `json:"message_globals"`
}

type MessageGlobals struct {
	Originator   string `json:"originator"`
	ReportURL    string `json:"report_url,omitempty"`
	ScheduleTime string `json:"schedule_time,omitempty"`
    Tag string `json:"tag,omitempty"`
}

type Message struct {
    Recipients []string `json:"recipients"`
    Content    string   `json:"content"`
    Unicode    string   `json:"unicode"`
}