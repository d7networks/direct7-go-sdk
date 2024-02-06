package direct7

import (
	"log"
)

// WhatsApp represents the WhatsApp service in Go.
type WhatsApp struct {
	client *Client
}

// NewWhatsApp creates a new instance of the WhatsApp service.
func NewWhatsApp(client *Client) *WhatsApp {
	return &WhatsApp{client: client}
}

type OptionalParams struct {
	firstName          string
	lastName           string
	displayName        string
	phone              string
	email              string
	url                string
	latitude           string
	longitude          string
	locationName       string
	locationAddress    string
	attachmentType     string
	attachmentURL      string
	attachmentCaption  string
	messageText        string
	mediaType	string
	mediaURL	string
	bodyParameterValues        map[string]interface{}

}

// SendWhatsAppFreeformMessage sends a WhatsApp message to one or more recipients.
func (w *WhatsApp) SendWhatsAppFreeformMessage(
	originator, recipient, messageType string, optParams *OptionalParams ) ([]byte, error) {
	message := map[string]interface{}{
		"originator": originator,
		"recipients": []map[string]interface{}{
			{"recipient": recipient},
		},
		"content": map[string]interface{}{
			"message_type": messageType,
		},
	}

	if optParams != nil {
		if optParams.firstName != "" || optParams.lastName != "" || optParams.displayName != "" || optParams.phone != "" || optParams.email != "" || optParams.url != "" {
			contact := map[string]interface{}{
				"first_name":  optParams.firstName,
				"last_name":   optParams.lastName,
				"display_name": optParams.displayName,
				"phone":      optParams.phone,
				"email":       optParams.email,
				"url":         optParams.url,
			}
			message["content"].(map[string]interface{})["contact"] = contact
		}

		if optParams.latitude != "" || optParams.longitude != "" || optParams.locationName != "" || optParams.locationAddress != "" {
			location := map[string]interface{}{
				"latitude":        optParams.latitude,
				"longitude":       optParams.longitude,
				"name":            optParams.locationName,
				"address":         optParams.locationAddress,
			}
			message["content"].(map[string]interface{})["location"] = location
		}

		if optParams.attachmentType != "" || optParams.attachmentURL != "" || optParams.attachmentCaption != "" {
			attachment := map[string]interface{}{
				"attachment_type":  optParams.attachmentType,
				"attachment_url":   optParams.attachmentURL,
				"attachment_caption": optParams.attachmentCaption,
			}
			message["content"].(map[string]interface{})["attachment"] = attachment
		}

		if optParams.messageText != "" {
			message["content"].(map[string]interface{})["message_text"] = optParams.messageText
		}
	}
	params := map[string]interface{}{
		"messages": []interface{}{message},
	}
	

	response, err := w.client.Post("/whatsapp/v1/send", true, params)
	if err != nil {
		return nil, err
	}
	log.Println("WhatsApp message sent successfully.")
	return response, nil
}

// SendWhatsAppTemplatedMessage sends a WhatsApp message using a template.
func (w *WhatsApp) SendWhatsAppTemplatedMessage(originator, recipient, templateID string, optParams *OptionalParams) ([]byte, error) {
	message := map[string]interface{}{
		"originator": originator,
		"recipients": []map[string]string{{"recipient": recipient}},
		"content": map[string]interface{}{
			"message_type": "TEMPLATE",
			"template": map[string]interface{}{
				"template_id":           templateID,
				"body_parameter_values": optParams.bodyParameterValues,
			},
		},
	}

	if optParams.mediaType != "" {
		if optParams.mediaType == "location" {
			message["content"].(map[string]interface{})["template"].(map[string]interface{})["media"] = map[string]interface{}{
				"media_type": optParams.mediaType,
				"location": map[string]string{
					"latitude":  optParams.latitude,
					"longitude": optParams.longitude,
					"name":      optParams.locationName,
					"address":   optParams.locationAddress,
				},
			}
		} else {
			message["content"].(map[string]interface{})["template"].(map[string]interface{})["media"] = map[string]interface{}{
				"media_type": optParams.mediaType,
				"media_url":  optParams.mediaURL,
			}
		}
	}

	response, err := w.client.Post("/whatsapp/v1/send", true, map[string]interface{}{"messages": []interface{}{message}})
	if err != nil {
		return nil, err
	}
	log.Println("Message sent successfully.")
	return response, nil
}

// GetStatus retrieves the status for a WhatsApp message request.
func (w *WhatsApp) GetStatus(requestID string) ([]byte, error) {
	response, err := w.client.Get("/whatsapp/v1/report/"+requestID, nil)
	if err != nil {
		return nil, err
	}
	log.Println("Message status retrieved successfully.")
	return response, nil
}
