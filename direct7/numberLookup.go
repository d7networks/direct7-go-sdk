package direct7

import (
	"log"
)

type NumberLookup struct {
	client *Client
}

func NewNumberLookup(client *Client) *NumberLookup {
	return &NumberLookup{client: client}
}

func (nl *NumberLookup) SearchNumberDetails(recipient string) (string, error) {
	params := map[string]interface {}{
		"recipient": recipient,
	}
	log.Println(params)
	response, err := nl.client.Post("/hlr/v1/lookup", true, params)
	if err != nil {
		return "", err
	}
	log.Println("Search request is successful.")
	return string(response), nil
}
