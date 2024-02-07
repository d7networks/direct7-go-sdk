package direct7

import (
	"log"
)

// NUMBER_LOOKUP struct represents the NUMBER_LOOKUP service in Go.
type NumberLookup struct {
	client *Client
}

// NewNumberLookup creates a new instance of the NUMBER_LOOKUP service.
func NewNumberLookup(client *Client) *NumberLookup {
	return &NumberLookup{client: client}
}

// SearchNumberDetails searches for number details.
func (nl *NumberLookup) SearchNumberDetails(recipient string) ([]byte, error) {
	params := map[string]interface {}{
		"recipient": recipient,
	}
	log.Println(params)
	response, err := nl.client.Post("/hlr/v1/lookup", true, params)
	if err != nil {
		return nil, err
	}
	log.Println("Search request is successful.")
	return response, nil
}
