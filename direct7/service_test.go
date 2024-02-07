package direct7

import (
	"testing"
)

func TestSendMessagesAndGetStatus(t *testing.T) {
	apiToken := "Your_Api_Token"
	client := NewClient(apiToken)
	// Define a test message
	testMessage := Message{
		Recipients: []string{"+919999XXXXXX"}, // Replace with the actual recipient number
		Content:    "Test message content",
		Unicode:    "false",
	}

	// Send the test message
	response, err := client.sms.SendMessages([]Message{testMessage}, "TestSender", "", "")
	if err != nil {
		t.Errorf("Failed to send message: %v", err)
		return
	}
	t.Logf("SendMessages response: %s", response)

	requestID := "001ff613-de30-4f82-81f6-1fe944b8f61b"
	// Get status of the sent message
	statusResponse, err := client.sms.GetStatus(requestID)
	if err != nil {
		t.Errorf("Failed to get message status: %v", err)
		return
	}
	t.Logf("GetStatus response: %s", statusResponse)
}

func TestNumberLookupSearchNumberDetails(t *testing.T) {
	apiToken := "Your_Api_Token"
	client := NewClient(apiToken)
	// Replace "RECIPIENT_PHONE_NUMBER" with the phone number you want to look up
	recipient := "+919999XXXXXX"

	response, err := client.numberLookup.SearchNumberDetails(recipient)
	if err != nil {
		t.Fatalf("Error searching number details: %v", err)
	}

	t.Logf("Number details search response: %s", response)
}

func TestSlackSendSlackMessage(t *testing.T) {
	apiToken := "Your_Api_Token"
	client := NewClient(apiToken)
	// Replace the placeholders with actual values
	content := "Test message content"
	workSpaceName := "D7-dev"
	channelName := "random"
	reportURL := "https://example.com/report"

	response, err := client.slack.SendSlackMessage(content, workSpaceName, channelName, reportURL)
	if err != nil {
		t.Fatalf("Error sending Slack message: %v", err)
	}

	t.Logf("Slack message sent successfully. Response: %s", response)
}

func TestSlackGetStatus(t *testing.T) {
	apiToken := "Your_Api_Token"
	client := NewClient(apiToken)
	// Replace "REQUEST_ID" with the actual request ID
	requestID := "001ff613-de30-4f82-81f6-1fe944b8f61b"

	response, err := client.slack.GetStatus(requestID)
	if err != nil {
		t.Fatalf("Error getting Slack message status: %v", err)
	}

	t.Logf("Slack message status retrieved successfully. Response: %s", response)
}

func TestVerifySendOTP(t *testing.T) {
	apiToken := "Your_Api_Token"
	client := NewClient(apiToken)
	// Replace these values with actual data
	originator := "SignOTP"
	recipient := "+919999XXXXXX"
	content := "Greetings from D7 API, your mobile verification code is: {}"
	dataCoding := "text"
	expiry := 600   // Expiry time in seconds
	templateID := 0 // Your verification template ID, if applicable

	response, err := client.verify.SendOTP(originator, recipient, content, dataCoding, expiry, templateID)
	if err != nil {
		t.Fatalf("Error sending OTP: %v", err)
	}

	t.Logf("OTP sent successfully. Response: %s", response)
}

func TestVerifyResendOTP(t *testing.T) {
	apiToken := "Your_Api_Token"
	client := NewClient(apiToken)
	// Replace "YOUR_OTP_ID" with the OTP ID you received
	otpID := "aeffa23f-1204-4e17-bb91-adf6de2cf826"

	response, err := client.verify.ResendOTP(otpID)
	if err != nil {
		t.Fatalf("Error resending OTP: %v", err)
	}

	t.Logf("OTP resent successfully. Response: %s", response)
}

func TestVerifyVerifyOTP(t *testing.T) {
	apiToken := "Your_Api_Token"
	client := NewClient(apiToken)
	// Replace "YOUR_OTP_ID" with the OTP ID you received
	otpID := "32549451-0eb6-4788-8e91-32b2eb9c4260"
	// Replace "YOUR_OTP_CODE" with the OTP code you received
	otpCode := "803053"

	response, err := client.verify.VerifyOTP(otpID, otpCode)
	if err != nil {
		t.Fatalf("Error verifying OTP: %v", err)
	}

	t.Logf("OTP verified successfully. Response: %s", response)
}

func TestVerifyGetStatus(t *testing.T) {
	apiToken := "Your_Api_Token"
	client := NewClient(apiToken)
	// Replace "YOUR_OTP_ID" with the OTP ID you received
	otpID := "32549451-0eb6-4788-8e91-32b2eb9c4260"

	response, err := client.verify.GetStatus(otpID)
	if err != nil {
		t.Fatalf("Error getting OTP status: %v", err)
	}

	t.Logf("OTP status retrieved successfully. Response: %s", response)
}

func TestViberSendViberMessage(t *testing.T) {
	apiToken := "Your_Api_Token"
	client := NewClient(apiToken)
	// Replace the values below with actual data
	recipients := []string{"+919999XXXXXX"}
	content := "Test Viber message"
	label := "PROMOTION"
	originator := "INFO2WAY"
	callBackURL := "https://example.com/callback"

	response, err := client.viber.SendViberMessage(recipients, content, label, originator, callBackURL)
	if err != nil {
		t.Fatalf("Error sending Viber message: %v", err)
	}

	t.Logf("Viber message sent successfully. Response: %s", response)
}

func TestViberGetStatus(t *testing.T) {
	apiToken := "Your_Api_Token"
	client := NewClient(apiToken)
	// Replace "YOUR_REQUEST_ID" with the actual request ID
	requestID := "642f4df4-a9a2-4be0-8834-d1f79f28c045"

	response, err := client.viber.GetStatus(requestID)
	if err != nil {
		t.Fatalf("Error getting Viber message status: %v", err)
	}

	t.Logf("Viber message status retrieved successfully. Response: %s", response)
}

func TestWhatsAppSendWhatsAppFreeformMessage(t *testing.T) {
	apiToken := "Your_Api_Token"
	client := NewClient(apiToken)
	// Replace with appropriate values for your test
	originator := "+91906XXXXXX"
	recipient := "+919999XXXXXX"
	messageType := "TEXT"
	optParams := &OptionalParams{messageText: "HI"}
	response, err := client.whatsapp.SendWhatsAppFreeformMessage(originator, recipient, messageType, optParams)
	if err != nil {
		t.Fatalf("Error sending WhatsApp message: %v", err)
	}

	t.Logf("WhatsApp message sent successfully. Response: %s", response)
}

func TestWhatsAppSendWhatsAppTemplatedMessage(t *testing.T) {
	apiToken := "Your_Api_Token"
	client := NewClient(apiToken)
	// Replace the placeholders with your test data
	originator := "+91906XXXXXX"
	recipient := "+919999XXXXXX"
	templateID := "marketing_media_image"
	optParams := &OptionalParams{mediaType: "image", mediaURL: "https://25428574.fs1.hubspotusercontent-eu1.net/hubfs/25428574/D7%20Logo%20rect.webp", bodyParameterValues: map[string]interface{}{
		"0": "Anu",
	}}

	response, err := client.whatsapp.SendWhatsAppTemplatedMessage(originator, recipient, templateID, optParams)
	if err != nil {
		t.Fatalf("Error sending WhatsApp message: %v", err)
	}

	t.Logf("WhatsApp message sent successfully. Response: %s", response)
}

func TestWhatsappGetStatus(t *testing.T) {
	apiToken := "Your_Api_Token"
	client := NewClient(apiToken)
	// Replace "YOUR_REQUEST_ID" with the actual request ID
	requestID := "4c055731-8bf3-4806-bed9-89a9046462a8"

	response, err := client.whatsapp.GetStatus(requestID)
	if err != nil {
		t.Fatalf("Error getting Whatsapp message status: %v", err)
	}

	t.Logf("Whatsapp message status retrieved successfully. Response: %s", response)
}