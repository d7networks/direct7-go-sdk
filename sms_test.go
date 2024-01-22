// sms_test.go
package direct7

import (
	"testing"
	"time"
)

func TestSendMessage(t *testing.T) {
	// Create a mock client or use a testing API key
	client := NewClient("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJhdXRoLWJhY2tlbmQ6YXBwIiwic3ViIjoiOTM2M2FmNTUtYWRmMS00Y2YzLWJhNjEtNGRjNWIxOTE4NGUwIn0.rctBTKBUO2FERmv_j75ItWACpUDQ7NG14v1PeXlM1ks", 30*time.Second, 10, 10, 3)
	sms := NewSMS(client)

	// Define test parameters
	recipients := []string{"+917306445534"}
	content := "Test message content"
	originator := "TestSender"
	reportURL := "http://example.com/report"
	unicode := false

	// Perform the test
	response, err := sms.SendMessage(recipients, content, originator, reportURL, unicode)

	// Check the result
	if err != nil {
		t.Errorf("Error sending SMS: %v", err)
	}

	// Add additional assertions based on the expected behavior of your API
	// For example, you might check the response content, status, etc.
	// Here, we are just checking if the response is not nil.
	if response == nil {
		t.Error("Expected a non-nil response")
	}
}

func TestGetStatus(t *testing.T) {
	// Create a mock client or use a testing API key
	client := NewClient("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJhdXRoLWJhY2tlbmQ6YXBwIiwic3ViIjoiOTM2M2FmNTUtYWRmMS00Y2YzLWJhNjEtNGRjNWIxOTE4NGUwIn0.rctBTKBUO2FERmv_j75ItWACpUDQ7NG14v1PeXlM1ks", 30*time.Second, 10, 10, 3)
	sms := NewSMS(client)

	// Define test parameters
	requestID := "test_request_id"

	// Perform the test
	response, err := sms.GetStatus(requestID)

	// Check the result
	if err != nil {
		t.Errorf("Error getting SMS status: %v", err)
	}

	// Add additional assertions based on the expected behavior of your API
	// For example, you might check the response content, status, etc.
	// Here, we are just checking if the response is not nil.
	if response == nil {
		t.Error("Expected a non-nil response")
	}
}
