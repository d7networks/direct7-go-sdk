package direct7

import (
	"testing"
)

func TestSendMessagesAndGetStatus(t *testing.T) {
	apiToken := "Your_Api_Token"
	client := NewClient(apiToken)
	testMessage := Message{
		Recipients: []string{"+919999XXXXXX"},
		Content:    "Test message content",
		Unicode:    "false",
	}

	response, err := client.sms.SendMessages([]Message{testMessage}, "TestSender", "", "")
	if err != nil {
		t.Errorf("Failed to send message: %v", err)
		return
	}
	t.Logf("SendMessages response: %s", response)

	requestID := "001ff613-de30-4f82-81f6-1fe944b8f61b"
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
	originator := "SignOTP"
	recipient := "+919999XXXXXX"
	content := "Greetings from D7 API, your mobile verification code is: {}"
	dataCoding := "text"
	expiry := 600
	templateID := 0

	response, err := client.verify.SendOTP(originator, recipient, content, dataCoding, expiry, templateID)
	if err != nil {
		t.Fatalf("Error sending OTP: %v", err)
	}

	t.Logf("OTP sent successfully. Response: %s", response)
}

func TestVerifyResendOTP(t *testing.T) {
	apiToken := "Your_Api_Token"
	client := NewClient(apiToken)
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
	otpID := "32549451-0eb6-4788-8e91-32b2eb9c4260"
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
	requestID := "642f4df4-a9a2-4be0-8834-d1f79f28c045"

	response, err := client.viber.GetStatus(requestID)
	if err != nil {
		t.Fatalf("Error getting Viber message status: %v", err)
	}

	t.Logf("Viber message status retrieved successfully. Response: %s", response)
}

func TestWhatsAppSendWhatsAppFreeformMessageNormal(t *testing.T) {
	apiToken := "Your_Api_Token"
	client := NewClient(apiToken)
	originator := "+91906XXXXXX"
	recipient := "+919999XXXXXX"
	messageType := "TEXT"
	optParams := &OptionalParams{body: "HI"}
	response, err := client.whatsapp.SendWhatsAppFreeformMessage(originator, recipient, messageType, optParams)
	if err != nil {
		t.Fatalf("Error sending WhatsApp message: %v", err)
	}

	t.Logf("WhatsApp message sent successfully. Response: %s", response)
}

func TestWhatsAppSendWhatsAppTemplatedMessage(t *testing.T) {
	apiToken := "Your_Api_Token"
	client := NewClient(apiToken)
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
	requestID := "4c055731-8bf3-4806-bed9-89a9046462a8"

	response, err := client.whatsapp.GetStatus(requestID)
	if err != nil {
		t.Fatalf("Error getting Whatsapp message status: %v", err)
	}

	t.Logf("Whatsapp message status retrieved successfully. Response: %s", response)
}

func TestWhatsAppSendWhatsAppFreeformMessage(t *testing.T) {
	apiToken := "API TOKEN"
	client := NewClient(apiToken)
	originator := "+XXXXXXXXXXXXXXXX"
	recipient := "XXXXXXXXXXXXXXXX"
	messageType := "CONTACTS"
	optParams := &OptionalParams{
		firstName:     "Amal",
		lastName:      "Anu",
		formattedName: "Amal Anu",
		phones:        []map[string]string{{"phone": "9181XXXXXXXX"}, {"phone": "9181XXXXXXXX"}},
		emails:        []map[string]string{{"email": "amal@gmail.com"}, {"email": "amal@gmail1.com"}},
	}
	response, err := client.whatsapp.SendWhatsAppFreeformMessage(originator, recipient, messageType, optParams)
	if err != nil {
		t.Fatalf("Error sending WhatsApp message: %v", err)
	}

	t.Logf("WhatsApp message sent successfully. Response: %s", response)
}

func TestWhatsAppSendWhatsAppFreeformMessage_Text(t *testing.T) {
	apiToken := "API TOKEN"
	client := NewClient(apiToken)
	originator := "+XXXXXXXXXXXXXXXX"
	recipient := "XXXXXXXXXXXXXXXX"
	messageType := "TEXT"
	optParams := &OptionalParams{
		body: "Hi",
	}
	response, err := client.whatsapp.SendWhatsAppFreeformMessage(originator, recipient, messageType, optParams)
	if err != nil {
		t.Fatalf("Error sending WhatsApp message: %v", err)
	}

	t.Logf("WhatsApp message sent successfully. Response: %s", response)
}

func TestWhatsAppSendWhatsAppFreeformMessage_Attachment(t *testing.T) {
	apiToken := "API TOKEN"
	client := NewClient(apiToken)
	originator := "+XXXXXXXXXXXXXXXX"
	recipient := "XXXXXXXXXXXXXXXX"
	messageType := "ATTACHMENT"
	optParams := &OptionalParams{
		attachmentType:    "image",
		attachmentURL:     "https://upload.wikimedia.org",
		attachmentCaption: "Test",
	}
	response, err := client.whatsapp.SendWhatsAppFreeformMessage(originator, recipient, messageType, optParams)
	if err != nil {
		t.Fatalf("Error sending WhatsApp message: %v", err)
	}

	t.Logf("WhatsApp message sent successfully. Response: %s", response)
}

func TestWhatsAppSendWhatsAppFreeformMessage_Location(t *testing.T) {
	apiToken := "API TOKEN"
	client := NewClient(apiToken)
	// client := NewWhatsApp(&Client{NewClient(apiToken)})
	originator := "+XXXXXXXXXXXXXXXX"
	recipient := "XXXXXXXXXXXXXXXX"
	messageType := "LOCATION"
	optParams := &OptionalParams{
		latitude:     "12.93803129081362",
		longitude:    "77.61088653615994",
		name:         "Karix Mobile Pvt Ltd",
		address:      "30, Hosur Rd, 7th Block, Koramangala, Bengaluru, Karnataka 560095",
	}
	response, err := client.whatsapp.SendWhatsAppFreeformMessage(originator, recipient, messageType, optParams)
	if err != nil {
		t.Fatalf("Error sending WhatsApp message: %v", err)
	}

	t.Logf("WhatsApp message sent successfully. Response: %s", response)
}

func TestWhatsAppSendWhatsAppTemplatedMessageLto(t *testing.T) {
	apiToken := "API TOKEN"
	client := NewClient(apiToken)
	originator := "+XXXXXXXXXXXXXXXX"
	recipient := "XXXXXXXXXXXXXXXX"
	templateID := "limited_time_offer"
	optParams := &OptionalParams{
		mediaURL:               "https://miro.medium.com/max/780/1*9Wdo1PuiJTZo0Du2A9JLQQ.jpeg",
		ltoExpirationTimeMS:    "1704272804",
		couponCode:             "DWS44",
	}
	response, err := client.whatsapp.SendWhatsAppTemplatedMessage(originator, recipient, templateID, optParams)
	if err != nil {
		t.Fatalf("Error sending WhatsApp message: %v", err)
	}

	t.Logf("WhatsApp message sent successfully. Response: %s", response)
}

func TestWhatsAppSendWhatsAppTemplatedMessageActions(t *testing.T) {
	apiToken := "API TOKEN"
	client := NewClient(apiToken)
	originator := "+XXXXXXXXXXXXXXXX"
	recipient := "XXXXXXXXXXXXXXXX"
	templateID := "click_me"
	bodyParameterValues:= map[string]interface{}{
    }
	actions := []map[string]interface{}{
		{
			"action_type":    "url",
			"action_index":   "0",
			"action_payload": "ButtonText",
		},
	}
	optParams := &OptionalParams{
		bodyParameterValues: bodyParameterValues,
		actions: actions,
	}
	response, err := client.whatsapp.SendWhatsAppTemplatedMessage(originator, recipient, templateID, optParams)
	if err != nil {
		t.Fatalf("Error sending WhatsApp message: %v", err)
	}

	t.Logf("WhatsApp message sent successfully. Response: %s", response)
}

func TestWhatsAppSendWhatsAppTemplatedMessageCarousel(t *testing.T) {
	apiToken := "API TOKEN"
	client := NewClient(apiToken)
	originator := "+XXXXXXXXXXXXXXXX"
	recipient := "XXXXXXXXXXXXXXXX"
	templateID := "carousel_card"
	cards := []map[string]interface{}{
		{
			"card_index": "0",
			"components": []map[string]interface{}{
				{
					"type": "header",
					"parameters": []map[string]interface{}{
						{
							"type": "image",
							"image": map[string]interface{}{
								"link": "https://miro.medium.com/max/780/1*9Wdo1PuiJTZo0Du2A9JLQQ.jpeg",
							},
						},
					},
				},
				{
					"type": "button",
					"sub_type": "quick_reply",
					"index": "0",
					"parameters": []map[string]interface{}{
						{
							"type": "payload",
							"payload": "2259NqSd",
						},
					},
				},
			},
		},
		{
			"card_index": "1",
			"components": []map[string]interface{}{
				{
					"type": "header",
					"parameters": []map[string]interface{}{
						{
							"type": "image",
							"image": map[string]interface{}{
								"link": "https://www.selfdrive.ae/banner_image/desktop/21112023164328_409449002729.jpg",
							},
						},
					},
				},
				{
					"type": "button",
					"sub_type": "quick_reply",
					"index": "0",
					"parameters": []map[string]interface{}{
						{
							"type": "payload",
							"payload": "59NqSdd",
						},
					},
				},
			},
		},
	}
	optParams := &OptionalParams{
		carousel_cards: cards,
	}
	response, err := client.whatsapp.SendWhatsAppTemplatedMessage(originator, recipient, templateID, optParams)
	if err != nil {
		t.Fatalf("Error sending WhatsApp message: %v", err)
	}

	t.Logf("WhatsApp message sent successfully. Response: %s", response)
}

