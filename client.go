package direct7

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

var JSONDecodeError = fmt.Errorf("JSON Decode Error")

type Client struct {
	apiToken       string
	host           string
	headers        map[string]string
	sms            SMS
	verify         Verify
	viber          Viber
	slack          Slack
	numberLookup   NumberLookup
	whatsapp       Whatsapp
	timeout        time.Duration
	session        *http.Client
}

func NewClient(apiToken string, timeout time.Duration, poolConnections, poolMaxSize, maxRetries int) *Client {
	client := &Client{
		apiToken: apiToken,
		host:     "https://api.d7networks.com",
		headers: map[string]string{
			"User-Agent": "direct7-go-sdk",
			"Accept":     "application/json",
		},
		sms:        NewSMS(),
		verify:     NewVerify(),
		viber:      NewViber(),
		slack:      NewSlack(),
		numberLookup: NewNumberLookup(),
		whatsapp:   NewWhatsapp(),
		timeout:    timeout,
		session: &http.Client{
			Transport: &http.Transport{
				MaxIdleConns:        poolConnections,
				MaxIdleConnsPerHost: poolMaxSize,
			},
			Timeout: timeout,
		},
	}
	client.session.Transport = &http.Transport{
		MaxIdleConns:        poolConnections,
		MaxIdleConnsPerHost: poolMaxSize,
	}

	return client
}

func (c *Client) createBearerTokenString() string {
	return "Bearer " + c.apiToken
}

func (c *Client) host(value ...string) string {
	if len(value) == 0 {
		return c.host
	}
	c.host = value[0]
	return c.host
}

func (c *Client) processResponse(host string, response *http.Response) (map[string]interface{}, error) {
	log.Printf("Response headers %v\n", response.Header)

	if response.StatusCode == http.StatusUnauthorized {
		return nil, fmt.Errorf("Invalid API token")
	} else if response.StatusCode >= 200 && response.StatusCode < 300 {
		// success response
		var result map[string]interface{}
		err := json.NewDecoder(response.Body).Decode(&result)
		if err != nil {
			return nil, err
		}
		log.Printf("Successful process response: %v\n", result)
		return result, nil
	} else if response.StatusCode >= 400 && response.StatusCode < 500 {
		log.Printf("Client error: %v %v\n", response.StatusCode, response.Body)
		if response.StatusCode == http.StatusBadRequest {
			return nil, fmt.Errorf("%v", response.Body)
		} else if response.StatusCode == http.StatusNotFound {
			return nil, fmt.Errorf("%v", response.Body)
		} else if response.StatusCode == http.StatusPaymentRequired {
			return nil, fmt.Errorf("%v", response.Body)
		} else if response.StatusCode == http.StatusUnprocessableEntity {
			return nil, fmt.Errorf("%v", response.Body)
		} else {
			return nil, fmt.Errorf("%v response from %v", response.StatusCode, host)
		}
	} else if response.StatusCode >= 500 && response.StatusCode < 600 {
		log.Printf("Server error: %v %v\n", response.StatusCode, response.Body)
		return nil, fmt.Errorf("%v response from %v", response.StatusCode, host)
	}

	return nil, fmt.Errorf("Unexpected response status code: %v", response.StatusCode)
}

func (c *Client) get(host, path string, params map[string]string) (map[string]interface{}, error) {
	requestURL := fmt.Sprintf("%v%v", host, path)
	c.headers["Authorization"] = c.createBearerTokenString()
	log.Printf("GET request sent to %v with headers %v and params %v\n", requestURL, c.headers, params)

	response, err := c.session.Get(requestURL)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	return c.processResponse(host, response)
}

func (c *Client) post(host, path string, bodyIsJSON bool, params map[string]interface{}) (map[string]interface{}, error) {
	requestURL := fmt.Sprintf("%v%v", host, path)
	c.headers["Authorization"] = c.createBearerTokenString()
	if bodyIsJSON {
		c.headers["Content-Type"] = "application/json"
		log.Printf("POST request sent to %v with headers %v and params %v\n", requestURL, c.headers, params)
		response, err := c.session.Post(requestURL, "application/json", nil)
		if err != nil {
			return nil, err
		}
		defer response.Body.Close()

		return c.processResponse(host, response)
	}

	c.headers["Content-Type"] = "application/x-www-form-urlencoded"
	log.Printf("POST request sent to %v with headers %v and params %v\n", requestURL, c.headers, params)
	response, err := c.session.Post(requestURL, "application/x-www-form-urlencoded", nil)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	return c.processResponse(host, response)
}

// Implement other HTTP methods (put, delete, patch) if needed

