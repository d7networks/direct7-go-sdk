package main

import (
	"log"
	"github.com/d7networks/direct7-go-sdk/direct7"
)

func main() {
	apiToken := "your-api-token"
	client := direct7.NewClient(apiToken)
	sms := direct7.NewSMS(client)

	// Example: Send SMS
	messages := []direct7.Message{ // Use the correct type from the direct7 package
		{
			Recipients: []string{"+918086757074"},
			Content:    "Test message 1",
			Unicode:    "false",
		},
	}

	originator := "Sender"
	reportURL := "https://example.com/report"
	scheduleTime := ""

	response, err := sms.SendMessages(messages, originator, reportURL, scheduleTime)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Response:", string(response))

	// Example: Get SMS status
	requestID := "your-request-id" // Replace with a valid request ID
	statusResponse, err := sms.GetStatus(requestID)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Status Response:", string(statusResponse))
}
