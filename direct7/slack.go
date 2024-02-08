package direct7

import (
	"fmt"
	"log"
)

type Slack struct {
	client *Client
}

func NewSlack(client *Client) *Slack {
	return &Slack{client: client}
}

func (s *Slack) SendSlackMessage(content, workSpaceName, channelName, reportURL string) (string, error) {
	message := map[string]interface{}{
		"channel":        "slack",
		"content":        content,
		"work_space_name": workSpaceName,
		"channel_name":   channelName,
	}
	messageGlobals := map[string]interface{}{
		"report_url": reportURL,
	}
	payload := map[string]interface{}{
		"messages":         []map[string]interface{}{message},
		"message_globals":  messageGlobals,
	}
	response, err := s.client.Post("/messages/v1/send", true, payload)
	if err != nil {
		return "", err
	}
	log.Println("Message sent successfully.")
	return string(response), nil
}

func (s *Slack) GetStatus(requestID string) (string, error) {
	response, err := s.client.Get(fmt.Sprintf("/report/v1/message-log/%s", requestID), nil)
	if err != nil {
		return "", err
	}
	log.Println("Message status retrieved successfully.")
	return string(response), nil
}
