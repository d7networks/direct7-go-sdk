package direct7

import (
	"log"
)

type WhatsApp struct {
	client *Client
}

func NewWhatsApp(client *Client) *WhatsApp {
	return &WhatsApp{client: client}
}

type OptionalParams struct {
	firstName           string
	lastName            string
	formattedName       string
	birthday            string
	phones              []map[string]string
	emails              []map[string]string
	urls                []map[string]string
	latitude            string
	longitude           string
	name                string
	address             string
	attachmentType      string
	attachmentURL       string
	attachmentCaption   string
	body                string
	mediaType           string
	mediaURL            string
	ltoExpirationTimeMS string
	couponCode          string
	bodyParameterValues map[string]interface{}
	actions             any
	quick_replies       any
	carousel_cards      any
}

func (w *WhatsApp) SendWhatsAppFreeformMessage(
	originator, recipient, messageType string, optParams *OptionalParams) (string, error) {
	message := map[string]interface{}{
		"originator": originator,
		"recipients": []map[string]interface{}{
			{"recipient": recipient},
		},
		"content": map[string]interface{}{
			"message_type": messageType,
		},
	}

	if messageType == "CONTACTS" {
		contacts := map[string]interface{}{
			"name": map[string]interface{}{
				"first_name":     optParams.firstName,
				"last_name":      optParams.lastName,
				"formatted_name": optParams.formattedName,
			},
			"birthday": optParams.birthday,
			"phones":   optParams.phones,
			"emails":   optParams.emails,
			"urls":     optParams.urls,
		}
		message["content"].(map[string]interface{})["contacts"] = []map[string]interface{}{contacts}
	} else if messageType == "LOCATION" {
		location := map[string]interface{}{
			"latitude":  optParams.latitude,
			"longitude": optParams.longitude,
			"name":      optParams.name,
			"address":   optParams.address,
		}
		message["content"].(map[string]interface{})["location"] = location
	} else if messageType == "ATTACHMENT" {
		attachment := map[string]interface{}{
			"type":    optParams.attachmentType,
			"url":     optParams.attachmentURL,
			"caption": optParams.attachmentCaption,
		}
		message["content"].(map[string]interface{})["attachment"] = attachment
	} else if messageType == "TEXT" {
		message["content"].(map[string]interface{})["text"] = map[string]interface{}{
			"body": optParams.body,
		}
	}

	response, err := w.client.Post("/whatsapp/v2/send", true, map[string]interface{}{"messages": []interface{}{message}})
	if err != nil {
		return "", err
	}
	log.Println("WhatsApp message sent successfully.")
	return string(response), nil
}

func (w *WhatsApp) SendWhatsAppTemplatedMessage(
	originator, recipient, templateID string, optParams *OptionalParams) (string, error) {
	message := map[string]interface{}{
		"originator": originator,
		"recipients": []map[string]interface{}{
			{"recipient": recipient},
		},
		"content": map[string]interface{}{
			"message_type": "TEMPLATE",
			"template": map[string]interface{}{
				"template_id":           templateID,
				"body_parameter_values": optParams.bodyParameterValues,
			},
		},
	}

	if optParams.mediaType == "location" {
		message["content"].(map[string]interface{})["template"].(map[string]interface{})["media"] = map[string]interface{}{
			"media_type": optParams.mediaType,
			"location": map[string]interface{}{
				"latitude":  optParams.latitude,
				"longitude": optParams.longitude,
				"name":      optParams.name,
				"address":   optParams.address,
			},
		}
	} else if optParams.mediaType != "" {
		message["content"].(map[string]interface{})["template"].(map[string]interface{})["media"] = map[string]interface{}{
			"media_type": optParams.mediaType,
			"media_url":  optParams.mediaURL,
		}
	}

	if optParams.ltoExpirationTimeMS != "" {
		message["content"].(map[string]interface{})["template"].(map[string]interface{})["limited_time_offer"] = map[string]interface{}{
			"expiration_time_ms": optParams.ltoExpirationTimeMS,
		}
	}

	if optParams.couponCode != "" {
		message["content"].(map[string]interface{})["template"].(map[string]interface{})["buttons"] = map[string]interface{}{
			"coupon_code": []map[string]interface{}{
				{
					"index":       0,
					"type":        "copy_code",
					"coupon_code": optParams.couponCode,
				},
			},
		}
	}

	if optParams.actions != "" {
		message["content"].(map[string]interface{})["template"].(map[string]interface{})["buttons"] = map[string]interface{}{
			"actions": optParams.actions,
		}
	}

	if optParams.quick_replies != "" {
		message["content"].(map[string]interface{})["template"].(map[string]interface{})["buttons"] = map[string]interface{}{
			"quick_replies": optParams.quick_replies,
		}
	}

	if optParams.carousel_cards != "" {
		message["content"].(map[string]interface{})["template"].(map[string]interface{})["carousel"] = map[string]interface{}{
			"cards": optParams.carousel_cards,
		}
	}

	response, err := w.client.Post("/whatsapp/v2/send", true, map[string]interface{}{"messages": []interface{}{message}})
	if err != nil {
		return "", err
	}
	log.Println("Message sent successfully.")
	return string(response), nil
}

func (w *WhatsApp) GetStatus(requestID string) (string, error) {
	response, err := w.client.Get("/whatsapp/v1/report/"+requestID, nil)
	if err != nil {
		return "", err
	}
	log.Println("Message status retrieved successfully.")
	return string(response), nil
}
