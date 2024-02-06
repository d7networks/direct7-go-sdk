package direct7

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
	"net/url"
	"log"
)

// Client struct represents the D7 API client in Go.
type Client struct {
	apiToken string
	host     string
	headers  http.Header
	sms      *SMS
	number_lookup	*NumberLookup
	// slack	*Slack
	// verify	*Verify
	// Add other services here
}

// NewClient creates a new instance of the D7 API client.
func NewClient(apiToken string) *Client {
	client := &Client{
		apiToken: apiToken,
		host:     "https://api.d7networks.com",
		headers: http.Header{
			"User-Agent": []string{"direct7-go-sdk"},
			"Accept":     []string{"application/json"},
		},
	}
	client.sms = NewSMS(client)
	client.number_lookup = NewNumberLookup(client)
	// client.slack = Slack(client)
	// client.verify = Verify(client)
	// Initialize other services here
	return client
}

// SetHost sets the host for the API client.
func (c *Client) SetHost(host string) {
	c.host = host
}

// SetHeader sets a custom header for the API client.
func (c *Client) SetHeader(key, value string) {
	c.headers.Set(key, value)
}

// processResponse handles the response from the D7 API.
func (c *Client) processResponse(response *http.Response) ([]byte, error) {
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	// Handle different HTTP status codes
	switch response.StatusCode {
	case http.StatusOK:
		var result map[string]interface{}
		err := json.Unmarshal(body, &result)
		if err != nil {
			return nil, err
		}
		return body, nil
	case http.StatusUnauthorized:
		return nil, errors.New("Invalid API token")
	case http.StatusBadRequest:
return nil, errors.New("Client error: " + fmt.Sprint(response.StatusCode) + " " + string(body))
	case http.StatusNotFound:
		return nil, errors.New("Not Found: " + string(body))
	case http.StatusPaymentRequired:
		return nil, errors.New("Insufficient Credit: " + string(body))
	case http.StatusUnprocessableEntity:
		return nil, errors.New("Validation Error: " + string(body))
	case http.StatusInternalServerError:
		return nil, errors.New("Server Error: " + string(body))
	default:
		return nil, errors.New("Unexpected response: " + string(body))
	}
}

// Get sends an HTTP GET request to the D7 API.
func (c *Client) Get(path string, params map[string]string) ([]byte, error) {
	url := c.host + path
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	request.Header = c.headers
	request.Header.Add("Authorization", "Bearer "+c.apiToken)

	query := request.URL.Query()
	for key, value := range params {
		query.Add(key, value)
	}
	request.URL.RawQuery = query.Encode()

	client := &http.Client{Timeout: 30 * time.Second}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	return c.processResponse(response)
}

func (c *Client) Post(path string, bodyIsJSON bool, payload interface{}) ([]byte, error) {
	apiURL := c.host + path
	var request *http.Request
	var err error
	if bodyIsJSON {
		jsonBody, err := json.Marshal(payload)
		if err != nil {
			return nil, err
		}

		request, err = http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonBody))
		if err != nil {
			return nil, err
		}
		request.Header.Set("Content-Type", "application/json")
	} else {
		formData := url.Values{}
		for key, value := range payload.(map[string]interface{}) {
			formData.Add(key, fmt.Sprintf("%v", value))
		}

		request, err = http.NewRequest("POST", apiURL, strings.NewReader(formData.Encode()))
		if err != nil {
			return nil, err
		}
		log.Println(strings.NewReader(formData.Encode()))
		request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}

	// Set headers for the request
	for key, values := range c.headers {
		for _, value := range values {
			request.Header.Add(key, value)
		}
	}
	request.Header.Add("Authorization", "Bearer "+c.apiToken)

	client := &http.Client{Timeout: 30 * time.Second}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	return c.processResponse(response)
}