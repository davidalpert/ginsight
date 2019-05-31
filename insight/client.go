package insight

import (
	"crypto/tls"
	"fmt"
	"log"
	"os"

	"gopkg.in/resty.v1"
)

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

// Applies ClientConfiguration to a Client, configuring the Resty client underneath
func (c *Client) Configure(config *ClientConfiguration) error {
	configErr := config.ValidateProperties()
	if configErr != nil {
		return configErr
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

// Exposes (or creates) a default insight.Client
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

func SetDefaultClient(c *Client) {
	defaultInsightClient = c
}

// Builds a new instance of an Insight API Client
func BuildClient(config *ClientConfiguration) (*Client, error) {
	configErr := config.ValidateProperties()
	if configErr != nil {
		return nil, configErr
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
