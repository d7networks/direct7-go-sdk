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

// Whatsapp: Text
func TestWhatsAppSendWhatsAppFreeformMessageText(t *testing.T) {
	apiToken := "Your_Api_Token"
	client := NewClient(apiToken)
	originator := "+97156XXXXXXXX"
	recipient := "+9180867XXXXXXX"
	messageType := "TEXT"
	optParams := &OptionalParams{body: "Hi"}
	response, err := client.whatsapp.SendWhatsAppFreeformMessage(originator, recipient, messageType, optParams)
	if err != nil {
		t.Fatalf("Error sending WhatsApp message: %v", err)
	}

	t.Logf("WhatsApp message sent successfully. Response: %s", response)
}

// Whatsapp: Reaction
func TestWhatsAppSendWhatsAppFreeformMessageText(t *testing.T) {
	apiToken := "Your_Api_Token"
	client := NewClient(apiToken)
	originator := "+97156XXXXXXXX"
	recipient := "+9180867XXXXXXX"
	messageType := "REACTION"
	optParams := &OptionalParams{messageId: "8916e1ac-26e2-11ef-acba-0242ac1b002e", emoji: "\U0001F600"}
	response, err := client.whatsapp.SendWhatsAppFreeformMessage(originator, recipient, messageType, optParams)
	if err != nil {
		t.Fatalf("Error sending WhatsApp message: %v", err)
	}

	t.Logf("WhatsApp message sent successfully. Response: %s", response)
}


// Whatsapp: Media: Image
func TestWhatsAppSendWhatsAppFreeformMessageText(t *testing.T) {
	apiToken := "Your_Api_Token"
	client := NewClient(apiToken)
	originator := "+97156XXXXXXXX"
	recipient := "+9180867XXXXXXX"
	messageType := "ATTACHMENT"
	optParams := &OptionalParams{
		attachmentType:"image",
		url: "https://miro.medium.com/max/780/1*9Wdo1PuiJTZo0Du2A9JLQQ.jpeg",
		caption: "Natural Beauty",
	}
	response, err := client.whatsapp.SendWhatsAppFreeformMessage(originator, recipient, messageType, optParams)
	if err != nil {
		t.Fatalf("Error sending WhatsApp message: %v", err)
	}

	t.Logf("WhatsApp message sent successfully. Response: %s", response)
}


// Whatsapp: Media: Video
func TestWhatsAppSendWhatsAppFreeformMessageText(t *testing.T) {
	apiToken := "Your_Api_Token"
	client := NewClient(apiToken)
	originator := "+97156XXXXXXXX"
	recipient := "+9180867XXXXXXX"
	messageType := "ATTACHMENT"
	optParams := &OptionalParams{
		attachmentType:"video",
		url: "http://www.onirikal.com/videos/mp4/nestlegold.mp4",
		caption: "D7 Video Test",
	}
	response, err := client.whatsapp.SendWhatsAppFreeformMessage(originator, recipient, messageType, optParams)
	if err != nil {
		t.Fatalf("Error sending WhatsApp message: %v", err)
	}

	t.Logf("WhatsApp message sent successfully. Response: %s", response)
}


// Whatsapp: Media: Documnet
func TestWhatsAppSendWhatsAppFreeformMessageText(t *testing.T) {
	apiToken := "Your_Api_Token"
	client := NewClient(apiToken)
	originator := "+97156XXXXXXXX"
	recipient := "+9180867XXXXXXX"
	messageType := "ATTACHMENT"
	optParams := &OptionalParams{
		attachmentType:"document",
		url: "https://www.clickdimensions.com/links/TestPDFfile.pdf",
		caption: "Test PDF file pdf",
		fileName: "TestPDFfile.pdf",
	}
	response, err := client.whatsapp.SendWhatsAppFreeformMessage(originator, recipient, messageType, optParams)
	if err != nil {
		t.Fatalf("Error sending WhatsApp message: %v", err)
	}

	t.Logf("WhatsApp message sent successfully. Response: %s", response)
}


// Whatsapp: Media: Audio
func TestWhatsAppSendWhatsAppFreeformMessageText(t *testing.T) {
	apiToken := "Your_Api_Token"
	client := NewClient(apiToken)
	originator := "+97156XXXXXXXX"
	recipient := "+9180867XXXXXXX"
	messageType := "ATTACHMENT"
	optParams := &OptionalParams{
		attachmentType:"audio",
		url: "http://fate-suite.ffmpeg.org/mpegaudio/extra_overread.mp3",
	}
	response, err := client.whatsapp.SendWhatsAppFreeformMessage(originator, recipient, messageType, optParams)
	if err != nil {
		t.Fatalf("Error sending WhatsApp message: %v", err)
	}

	t.Logf("WhatsApp message sent successfully. Response: %s", response)
}

// Whatsapp: Media: Sticker
func TestWhatsAppSendWhatsAppFreeformMessageText(t *testing.T) {
	apiToken := "Your_Api_Token"
	client := NewClient(apiToken)
	originator := "+97156XXXXXXXX"
	recipient := "+9180867XXXXXXX"
	messageType := "ATTACHMENT"
	optParams := &OptionalParams{
		attachmentType:"sticker",
		url: "https://raw.githubusercontent.com/sagarbhavsar4328/dummys3bucket/master/sample3.webp",
	}
	response, err := client.whatsapp.SendWhatsAppFreeformMessage(originator, recipient, messageType, optParams)
	if err != nil {
		t.Fatalf("Error sending WhatsApp message: %v", err)
	}

	t.Logf("WhatsApp message sent successfully. Response: %s", response)
}

// Whatsapp: Location
func TestWhatsAppSendWhatsAppFreeformMessageText(t *testing.T) {
	apiToken := "Your_Api_Token"
	client := NewClient(apiToken)
	originator := "+97156XXXXXXXX"
	recipient := "+9180867XXXXXXX"
	messageType := "LOCATION"
	optParams := &OptionalParams{
		longitude: "77.1249",
		latitude: "12.8779",
		name: "NameOfLocation",
		address: "AddressOfLocation",
	}
	response, err := client.whatsapp.SendWhatsAppFreeformMessage(originator, recipient, messageType, optParams)
	if err != nil {
		t.Fatalf("Error sending WhatsApp message: %v", err)
	}

	t.Logf("WhatsApp message sent successfully. Response: %s", response)
}


// Whatsapp: Contacts
func TestWhatsAppSendWhatsAppFreeformMessageText(t *testing.T) {
	apiToken := "Your_Api_Token"
	client := NewClient(apiToken)
	originator := "+97156XXXXXXXX"
	recipient := "+9180867XXXXXXX"
	messageType := "CONTACTS"
	contactAddresses := []map[string]string{
        {
            "Street":      "1 Lucky Shrub Way",
            "City":        "Menlo Park",
            "State":       "CA",
            "Zip":         "94025",
            "Country":     "United States",
            "CountryCode": "US",
            "Type":        "WORK",
        },
        {
            "Street":      "1 Hacker Way",
            "City":        "Menlo Park",
            "State":       "CA",
            "Zip":         "94025",
            "Country":     "United States",
            "CountryCode": "US",
            "Type":        "WORK",
        },
    }
    phones := []map[string]string{
        {
            "Phone": "+16505559999",
            "Type":  "HOME",
        },
        {
            "Phone": "+19175559999",
            "Type":  "WORK",
            "WaID":  "19175559999",
        },
    }
    emails := []map[string]string{
        {
            "Email": "bjohnson@luckyshrub.com",
            "Type":  "WORK",
        },
        {
            "Email": "bjohnson@luckyshrubplants.com",
            "Type":  "HOME",
        },
    }
    urls := []map[string]string{
        {
            "URL":  "https://www.luckyshrub.com",
            "Type": "WORK",
        },
        {
            "URL":  "https://www.facebook.com/luckyshrubplants",
            "Type": "WORK",
        },
    }
    optParams := &OptionalParams{
        firstName:        "Alice",
        lastName:         "Jane",
        formattedName:    "Alice Jane",
        middleName:       "Joana",
        suffix:           "Esq.",
        prefix:           "Dr.",
        phones:           phones,
        emails:           emails,
        contactAddresses: contactAddresses,
        urls:             urls,
    }
	response, err := client.whatsapp.SendWhatsAppFreeformMessage(originator, recipient, messageType, optParams)
	if err != nil {
		t.Fatalf("Error sending WhatsApp message: %v", err)
	}

	t.Logf("WhatsApp message sent successfully. Response: %s", response)
}


// Whatsapp: Interactive: cta
func TestWhatsAppSendWhatsAppFreeformMessageText(t *testing.T) {
	apiToken := "Your_Api_Token"
	client := NewClient(apiToken)
	originator := "+97156XXXXXXXX"
	recipient := "+9180867XXXXXXX"
	interactiveType := "cta_url"
	parameters := map[string]interface{}{
        "display_text": "Visit Us",
		"url": "https://www.luckyshrub.com?clickID=kqDGWd24Q5TRwoEQTICY7W1JKoXvaZOXWAS7h1P76s0R7Paec4",
    }
	optParams := &OptionalParams{
		headerType: "text",
		headerText: "Payment$ for D7 Whatsapp Service",
		bodyText: "Direct7 Networks is a messaging service provider that specializes in helping organizations efficiently communicate with their customers.",
		footerText: "Thank You",
		parameters: parameters,
	}
	response, err := client.whatsapp.SendWhatsAppInteractiveMessage(originator, recipient, interactiveType, optParams)
	if err != nil {
		t.Fatalf("Error sending WhatsApp message: %v", err)
	}

	t.Logf("WhatsApp message sent successfully. Response: %s", response)
}

// WhatsApp :Interactive flow
func TestWhatsAppSendWhatsAppFreeformMessageText(t *testing.T) {
	apiToken := "Your_Api_Token"
	client := NewClient(apiToken)
	originator := "+97156XXXXXXXX"
	recipient := "+9180867XXXXXXX"
	interactiveType := "flow"
	flowParameter := map[string]interface{}{
        "name": "flow",
        "parameters": map[string]interface{}{
            "flow_message_version": "3",
            "flow_token":           "unused",
            "flow_id":              "530404409952136",
            "flow_cta":             "Book Demo",
            "flow_action":          "navigate",
            "flow_action_payload": map[string]interface{}{
                "screen": "screen_",
            },
        },
    }
	 optParams := &OptionalParams{
        headerType: "text",
        headerText: "Payment$ for D7 Whatsapp Service",
        bodyText: "Direct7 Networks is a messaging service provider that specializes in helping organizations efficiently communicate with their customers.",
        footerText: "Thank You",
        parameters: flowParameter,
    }
    response, err := client.whatsapp.SendWhatsAppInteractiveMessage(originator, recipient, interactiveType, optParams)
	if err != nil {
		t.Fatalf("Error sending WhatsApp message: %v", err)
	}

	t.Logf("WhatsApp message sent successfully. Response: %s", response)
}
// Whatsapp: Interactive: button
func TestWhatsAppSendWhatsAppFreeformMessageText(t *testing.T) {
	apiToken := "Your_Api_Token"
	client := NewClient(apiToken)
	originator := "+97156XXXXXXXX"
	recipient := "+9180867XXXXXXX"
	interactiveType := "list"
	sections := []map[string]interface{}{
		{
			"title": "SMS Messaging",
			"rows": []map[string]interface{}{
				{
					"id":          "1",
					"title":       "Normal SMS",
					"description": "Signup for free at the D7 platform to use our Messaging APIs.",
				},
				{
					"id":          "2",
					"title":       "Verify",
					"description": "D7 Verify API is to applications requires SMS based OTP authentications.",
				},
			},
		},
		{
			"title": "Whatsapp",
			"rows": []map[string]interface{}{
				{
					"id":          "3",
					"title":       "WhatsApp Messages",
					"description": "D7 Whatsapp API is to applications requires pre-registration.",
				},
			},
		},
	}
	optParams := &OptionalParams{
		headerType: "text",
		headerText: "Payment$ for D7 Whatsapp Service",
		bodyText: "Direct7 Networks is a messaging service provider that specializes in helping organizations efficiently communicate with their customers.",
		footerText: "Thank You",
		sections: sections,
		listButtonText: "Choose Service",
	}
	response, err := client.whatsapp.SendWhatsAppInteractiveMessage(originator, recipient, interactiveType, optParams)
	if err != nil {
		t.Fatalf("Error sending WhatsApp message: %v", err)
	}

	t.Logf("WhatsApp message sent successfully. Response: %s", response)
}


// // Whatsapp: // Templated: no body parm
func TestWhatsAppSendWhatsAppFreeformMessageText(t *testing.T) {
	apiToken := "Your_Api_Token"
	client := NewClient(apiToken)
	originator := "+97156XXXXXXXX"
	recipient := "+9180867XXXXXXX"
	templateId := "testing_alpha"
	language := "en"
	optParams := &OptionalParams{
	}
	response, err := client.whatsapp.SendWhatsAppTemplatedMessage(originator, recipient, templateId, language, optParams)
	if err != nil {
		t.Fatalf("Error sending WhatsApp message: %v", err)
	}

	t.Logf("WhatsApp message sent successfully. Response: %s", response)
}


// Whatsapp: // Templated: with body parm
func TestWhatsAppSendWhatsAppFreeformMessageText(t *testing.T) {
	apiToken := "Your_Api_Token"
	client := NewClient(apiToken)
	originator := "+97156XXXXXXXX"
	recipient := "+9180867XXXXXXX"
	templateId := "with_personalize"
	language := "en"
	optParams := &OptionalParams{bodyParameterValues: map[string]interface{}{
				"0": "Anu",
	}}
	response, err := client.whatsapp.SendWhatsAppTemplatedMessage(originator, recipient, templateId, language, optParams)
	if err != nil {
		t.Fatalf("Error sending WhatsApp message: %v", err)
	}

	t.Logf("WhatsApp message sent successfully. Response: %s", response)
}


// // Whatsapp: Templated: media: text
func TestWhatsAppSendWhatsAppFreeformMessageText(t *testing.T) {
	apiToken := "Your_Api_Token"
	client := NewClient(apiToken)
	originator := "+97156XXXXXXXX"
	recipient := "+9180867XXXXXXX"
	templateId := "header_param"
	language := "en"
	optParams := &OptionalParams{textHeaderTitle: "Tom"}
	response, err := client.whatsapp.SendWhatsAppTemplatedMessage(originator, recipient, templateId, language, optParams)
	if err != nil {
		t.Fatalf("Error sending WhatsApp message: %v", err)
	}

	t.Logf("WhatsApp message sent successfully. Response: %s", response)
}


// // Whatsapp: Templated: media: image
func TestWhatsAppSendWhatsAppFreeformMessageText(t *testing.T) {
	apiToken := "Your_Api_Token"
	client := NewClient(apiToken)
	originator := "+97156XXXXXXXX"
	recipient := "+9180867XXXXXXX"
	templateId := "image"
	language := "en"
	optParams := &OptionalParams{mediaType: "image", mediaURL: "https://miro.medium.com/max/780/1*9Wdo1PuiJTZo0Du2A9JLQQ.jpeg"}
	response, err := client.whatsapp.SendWhatsAppTemplatedMessage(originator, recipient, templateId, language, optParams)
	if err != nil {
		t.Fatalf("Error sending WhatsApp message: %v", err)
	}

	t.Logf("WhatsApp message sent successfully. Response: %s", response)
}


// // Whatsapp: Templated: media: video
func TestWhatsAppSendWhatsAppFreeformMessageText(t *testing.T) {
	apiToken := "Your_Api_Token"
	client := NewClient(apiToken)
	originator := "+97156XXXXXXXX"
	recipient := "+9180867XXXXXXX"
	templateId := "video"
	language := "en"
	optParams := &OptionalParams{mediaType: "video", mediaURL: "http://www.onirikal.com/videos/mp4/nestlegold.mp4"}
	response, err := client.whatsapp.SendWhatsAppTemplatedMessage(originator, recipient, templateId, language, optParams)
	if err != nil {
		t.Fatalf("Error sending WhatsApp message: %v", err)
	}

	t.Logf("WhatsApp message sent successfully. Response: %s", response)
}

// // Whatsapp: Templated: media: document
func TestWhatsAppSendWhatsAppFreeformMessageText(t *testing.T) {
	apiToken := "Your_Api_Token"
	client := NewClient(apiToken)
	originator := "+97156XXXXXXXX"
	recipient := "+9180867XXXXXXX"
	templateId := "document"
	language := "en"
	optParams := &OptionalParams{
		bodyParameterValues: map[string]interface{}{ "0": "Anu", },
		mediaType: "document", 
		mediaURL: "https://www.clickdimensions.com/links/TestPDFfile.pdf"}
	response, err := client.whatsapp.SendWhatsAppTemplatedMessage(originator, recipient, templateId, language, optParams)
	if err != nil {
		t.Fatalf("Error sending WhatsApp message: %v", err)
	}

	t.Logf("WhatsApp message sent successfully. Response: %s", response)
}


// // Whatsapp: Templated: media: location
func TestWhatsAppSendWhatsAppFreeformMessageText(t *testing.T) {
	apiToken := "Your_Api_Token"
	client := NewClient(apiToken)
	originator := "+97156XXXXXXXX"
	recipient := "+9180867XXXXXXX"
	templateId := "location"
	language := "en"
	optParams := &OptionalParams{
		mediaType: "location", 
		latitude : "12.93803129081362", 
		longitude : "77.61088653615994",
		name : "Mobile Pvt Ltd", 
		address : "Bengaluru, Karnataka 560095",
	}
	response, err := client.whatsapp.SendWhatsAppTemplatedMessage(originator, recipient, templateId, language, optParams)
	if err != nil {
		t.Fatalf("Error sending WhatsApp message: %v", err)
	}

	t.Logf("WhatsApp message sent successfully. Response: %s", response)
}


// // Whatsapp: Templated: quick_replies
func TestWhatsAppSendWhatsAppFreeformMessageText(t *testing.T) {
	apiToken := "Your_Api_Token"
	client := NewClient(apiToken)
	originator := "+97156XXXXXXXX"
	recipient := "+9180867XXXXXXX"
	templateId := "quick_reply"
	language := "en"
	quickReplies := []map[string]interface{}{
		{
			"button_index": "0",
			"button_payload": "1",
		},
		{
			"button_index": "1",
			"button_payload": "2",
		},
	}
	optParams := &OptionalParams{
		quickReplies: quickReplies,
	}
	response, err := client.whatsapp.SendWhatsAppTemplatedMessage(originator, recipient, templateId, language, optParams)
	if err != nil {
		t.Fatalf("Error sending WhatsApp message: %v", err)
	}

	t.Logf("WhatsApp message sent successfully. Response: %s", response)
}
//    WhatsApp :Templated :button_flow
func TestWhatsAppSendWhatsAppFreeformMessageText(t *testing.T) {
	apiToken := "Your_Api_Token"
	client := NewClient(apiToken)
	originator := "+97156XXXXXXXX"
	recipient := "+9180867XXXXXXX"
	templateId := "call_to_action"
	language := "en"
	buttonFlow := []map[string]interface{}{
		{
			"flow_token":     "unused",
			"action_type":    "flow",
			"index":          "0",
			"flow_action_data": map[string]interface{}{},
		},
	}
	optParams := &OptionalParams{
		ButtonFlow: buttonFlow,
	}
	response, err := client.whatsapp.SendWhatsAppTemplatedMessage(originator, recipient, templateId, language, optParams)
	if err != nil {
		t.Fatalf("Error sending WhatsApp message: %v", err)
	}

// // Whatsapp: Templated: actions
func TestWhatsAppSendWhatsAppFreeformMessageText(t *testing.T) {
	apiToken := "Your_Api_Token"
	client := NewClient(apiToken)
	originator := "+97156XXXXXXXX"
	recipient := "+9180867XXXXXXX"
	templateId := "call_to_action"
	language := "en"
	actions := []map[string]interface{}{
		{
			"action_type": "url",
			"action_index": "0",
			"action_payload": "dash",
		},
	}
	optParams := &OptionalParams{
		actions: actions,
	}
	response, err := client.whatsapp.SendWhatsAppTemplatedMessage(originator, recipient, templateId, language, optParams)
	if err != nil {
		t.Fatalf("Error sending WhatsApp message: %v", err)
	}

	t.Logf("WhatsApp message sent successfully. Response: %s", response)
}


// Whatsapp: Templated: coupon_code
func TestWhatsAppSendWhatsAppFreeformMessageText(t *testing.T) {
	apiToken := "Your_Api_Token"
	client := NewClient(apiToken)
	originator := "+97156XXXXXXXX"
	recipient := "+9180867XXXXXXX"
	templateId := "coupon_code"
	language := "en"
	optParams := &OptionalParams{bodyParameterValues: map[string]interface{}{
				"0": "Anu",},
				couponCode: "DER556"}

	response, err := client.whatsapp.SendWhatsAppTemplatedMessage(originator, recipient, templateId, language, optParams)
	if err != nil {
		t.Fatalf("Error sending WhatsApp message: %v", err)
	}

	t.Logf("WhatsApp message sent successfully. Response: %s", response)
}

// Whatsapp: Templated: LTO
func TestWhatsAppSendWhatsAppFreeformMessageText(t *testing.T) {
	apiToken := "Your_Api_Token"
	client := NewClient(apiToken)
	originator := "+97156XXXXXXXX"
	recipient := "+9180867XXXXXXX"
	templateId := "coupon_code"
	language := "en"
	optParams := &OptionalParams{
		mediaURL:               "https://miro.medium.com/max/780/1*9Wdo1PuiJTZo0Du2A9JLQQ.jpeg",
		ltoExpirationTimeMS:    "1704272804",
		couponCode:             "DWS44",
	}

	response, err := client.whatsapp.SendWhatsAppTemplatedMessage(originator, recipient, templateId, language, optParams)
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

