package insight

import (
	"crypto/tls"
	"fmt"
	"log"
	"os"

	resty "gopkg.in/resty.v1"
)

// Client represents a Jira Insight API client
type Client struct {
	BaseURL  string
	Username string
	Debug    bool
	Insecure bool

	// this includes an implied Client.Client property of
	// type *restyClient but also promotes all the Resty
	// *Client methods to be sent directly to values of
	// this struct.
	*resty.Client
}

// Configure applies a ClientConfiguration to a Client, also configuring the Resty client underneath
func (c *Client) Configure(config *ClientConfiguration) error {
	err := config.ValidateProperties()
	if err != nil {
		return err
	}

	// ensure that our client's Client property has a value
	if c.Client == nil {
		c.Client = resty.New()
	}

	// grab the underlying resty.Client and configure it
	restyClient := c.Client

	// https://stackoverflow.com/a/48523194/8997
	if config.Insecure {
		restyClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	}
	restyClient.SetHeader("Accept", "application/json")
	restyClient.SetHeader("Content-Type", "application/json")
	restyClient.SetBasicAuth(config.Username, config.Password)
	restyClient.SetRESTMode()
	restyClient.SetDebug(config.Debug)

	// now push the configuration into our client properties to reflect that
	c.BaseURL = config.BaseURL
	c.Username = config.Username
	c.Debug = config.Debug
	c.Insecure = config.Insecure

	return nil
}

var defaultInsightClient *Client

// DefaultClient exposes (or creates) a default insight.Client
func DefaultClient() *Client {
	if defaultInsightClient == nil {
		client, err := BuildClient(DefaultClientConfiguration())
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defaultInsightClient = client
	}

	return defaultInsightClient
}

// SetDefaultClient allows setting the default insight.Client if needed
func SetDefaultClient(c *Client) {
	defaultInsightClient = c
}

// BuildClient builds a new instance of an Insight API Client
func BuildClient(config *ClientConfiguration) (*Client, error) {
	err := config.ValidateProperties()
	if err != nil {
		return nil, err
	}

	if config.Debug {
		log.Printf("creating an insight api client pointing at %s", config.BaseURL)
	}

	insightClient := &Client{}
	insightClient.Configure(config)

	if config.Debug {
		log.Printf(" created an insight api client pointing at %s", config.BaseURL)
	}

	return insightClient, nil
}

func validateResponseCodeExact(response *resty.Response, expectedStatusCode int) error {
	statusCode := response.StatusCode()
	if statusCode != expectedStatusCode {
		return &ClientError{Response: response}
	}
	return nil
}

func validateResponseCodeInRange(response *resty.Response, lowInclusive int, highExclusive int) error {
	statusCode := response.StatusCode()
	if statusCode < lowInclusive || highExclusive <= statusCode {
		return &ClientError{Response: response}
	}
	return nil
}

/*
func getAndValidate(req *resty.Request, url string, lowInclusive int, highExclusive int) (*resty.Response, error) {
	response, err := req.Get(url)
	if err != nil {
		log.Printf("Error: %v\n", err)
		return nil, err
	}

	statusCode := response.StatusCode()
	if statusCode < lowInclusive || highExclusive <= statusCode {
		clientError := &ClientError{Response: response}
		return nil, clientError
	}

	return response, nil
}
*/
