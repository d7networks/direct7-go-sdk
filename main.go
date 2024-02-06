package main

import (
	"log"
	"github.com/d7networks/direct7-go-sdk"
)

func main() {
	apiToken := "Your Api Token"
	client := direct7.NewClient(apiToken)

	// Example: Send SMS
	messages := []Message{
		{
			recipients:  []string{"+919999XXXXXX"},
			content:     "Test message 1",
			unicode:  "false",
		}
		// Add more test messages as needed
	}

	originator := "Sender"
	reportURL := "https://example.com/report"
	scheduleTime := ""

	response, err := client.SMS.SendMessages(messages, originator, reportURL, scheduleTime)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Response:", string(response))

	// Example: Get SMS status
	requestID := "your-request-id"  // Replace with a valid request ID
	statusResponse, err := client.SMS.GetStatus(requestID)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Status Response:", string(statusResponse))
}
