package main

import (
	"log"
	"github.com/d7networks/direct7-go-sdk"
)

func main() {
	apiToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJhdXRoLWJhY2tlbmQ6YXBwIiwic3ViIjoiOTM2M2FmNTUtYWRmMS00Y2YzLWJhNjEtNGRjNWIxOTE4NGUwIn0.rctBTKBUO2FERmv_j75ItWACpUDQ7NG14v1PeXlM1ks"
	client := direct7.NewClient(apiToken)

	// Example: Send SMS
	messages := []Message{
		{
			recipients:  []string{"+918086757074"},
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
