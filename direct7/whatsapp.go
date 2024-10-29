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
	middleName       	string
	suffix				string
	prefix				string
	birthday            string
	phones              []map[string]string
	emails              []map[string]string
	urls                []map[string]string
	contactAddresses   	[]map[string]string
	latitude            string
	longitude           string
	name                string
	address             string
	attachmentType		string
	url       			string
	caption   			string
	fileName   			string
	body                string
	messageId			string
	emoji				string
	mediaType           string
	mediaURL            string
	ltoExpirationTimeMS string
	couponCode          string
	bodyParameterValues map[string]interface{}
	textHeaderTitle 	string
	actions             any
	quickReplies       	any
	carouselCards      	any
	headerType 			string
	headerText 			string
	headerLink 			string
	headerFileName 		string
	bodyText 			string
	footerText 			string
	parameters 			map[string]interface{}
	sections 			[]map[string]interface{}
	buttons 			[]map[string]interface{}
	listButtonText 		string
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
				"middle_name": optParams.middleName,
				"suffix": optParams.suffix,
				"prefix": optParams.prefix,
			},
			"birthday": optParams.birthday,
			"phones":   optParams.phones,
			"emails":   optParams.emails,
			"urls":     optParams.urls,
			"addresses": optParams.contactAddresses,
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
		if optParams.attachmentType == "document" {
			attachment := map[string]interface{}{
				"type":     optParams.attachmentType,
				"url":      optParams.url,
				"caption":  optParams.caption,
				"filename": optParams.fileName,
			}
			message["content"].(map[string]interface{})["attachment"] = attachment
		} else {
			attachment := map[string]interface{}{
				"type":    optParams.attachmentType,
				"url":     optParams.url,
				"caption": optParams.caption,
			}
			message["content"].(map[string]interface{})["attachment"] = attachment
		}
	} else if messageType == "TEXT" {
		message["content"].(map[string]interface{})["text"] = map[string]interface{}{
			"body": optParams.body,
		}
	} else if messageType == "REACTION" {
		message["content"].(map[string]interface{})["reaction"] = map[string]interface{}{
			"message_id": optParams.messageId,
			"emoji": optParams.emoji,
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
	originator, recipient, templateId string, language string, optParams *OptionalParams) (string, error) {
	message := map[string]interface{}{
		"originator": originator,
		"recipients": []map[string]interface{}{
			{"recipient": recipient},
		},
		"content": map[string]interface{}{
			"message_type": "TEMPLATE",
			"template": map[string]interface{}{
				"template_id":	templateId,
				"language":	language,	
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
	} else if optParams.mediaType == "text" {
		message["content"].(map[string]interface{})["template"].(map[string]interface{})["media"] = map[string]interface{}{
			"media_type": optParams.mediaType,
			"text_header_title":  optParams.textHeaderTitle,
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

	if optParams.quickReplies != "" {
		message["content"].(map[string]interface{})["template"].(map[string]interface{})["buttons"] = map[string]interface{}{
			"quickReplies": optParams.quickReplies,
		}
	}

	if optParams.carouselCards != nil {
		message["content"].(map[string]interface{})["template"].(map[string]interface{})["carousel"] = map[string]interface{}{
			"cards": optParams.carouselCards,
		}
	}

	response, err := w.client.Post("/whatsapp/v2/send", true, map[string]interface{}{"messages": []interface{}{message}})
	if err != nil {
		return "", err
	}
	log.Println("Message sent successfully.")
	return string(response), nil
}


func (w *WhatsApp) SendWhatsAppInteractiveMessage(
	originator, recipient, interactiveType string, optParams *OptionalParams) (string, error) {
	message := map[string]interface{}{
		"originator": originator,
		"recipients": []map[string]interface{}{
			{"recipient": recipient},
		},
		"content": map[string]interface{}{
			"message_type": "INTERACTIVE",
			"interactive": map[string]interface{}{
				"type":	interactiveType,
				"header": map[string]interface{}{
					"type": optParams.headerType,
				},
				"body": map[string]interface{}{
					"text": optParams.bodyText,
				},
				"footer": map[string]interface{}{
					"text": optParams.footerText,
				},
			},
		},
	}

	if optParams.headerType == "text" {
		message["content"].(map[string]interface{})["interactive"].(map[string]interface{})["header"].(map[string]interface{})["text"] = optParams.headerText
	} else if optParams.headerType == "image" || optParams.headerType == "video" || optParams.headerType == "document" {
		headerContent := map[string]interface{}{
			"link": optParams.headerLink,
		}
		if optParams.headerType == "document" {
			headerContent["filename"] = optParams.headerFileName
		}
		message["content"].(map[string]interface{})["interactive"].(map[string]interface{})["header"].(map[string]interface{})[optParams.headerType] = headerContent
	}
	if interactiveType == "cta_url" {
		message["content"].(map[string]interface{})["interactive"].(map[string]interface{})["action"] = map[string]interface{}{
			"parameters": optParams.parameters,
		}
	} else if interactiveType == "button" {
		message["content"].(map[string]interface{})["interactive"].(map[string]interface{})["action"] = map[string]interface{}{
			"buttons": optParams.buttons,
		}
	} else if interactiveType == "list" {
		message["content"].(map[string]interface{})["interactive"].(map[string]interface{})["action"] = map[string]interface{}{
			"sections": optParams.sections,
			"button": optParams.listButtonText,
		}
	} else if interactiveType == "location_request_message" {
		message["content"].(map[string]interface{})["interactive"].(map[string]interface{})["action"] = map[string]interface{}{
			"name": "send_location",
		}
	} else if interactiveType == "address_message" {
		message["content"].(map[string]interface{})["interactive"].(map[string]interface{})["action"] = map[string]interface{}{
			"parameters": optParams.parameters,
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
