package api

import (
	"fmt"

	"github.com/spf13/viper"
)

// ClientConfiguration provides properties to configure an api.Client
type ClientConfiguration struct {
	BaseURL  string
	Username string
	Password string
	Debug    bool
	Insecure bool
}

// ValidateProperties validates a ClientConfiguration, creating a new value if the pointer is nil
func (c *ClientConfiguration) ValidateProperties() error {
	if c == nil {
		// set the config pointer to the address of a newly
		// created struct initialized with default values
		c = DefaultClientConfiguration()
	}

	if c.BaseURL == "" {
		return fmt.Errorf("must provide a baseurl\n")
	}
	if c.Username == "" {
		return fmt.Errorf("must provide a username\n")
	}
	if c.Password == "" {
		return fmt.Errorf("must provide a password\n")
	}

	return nil
}

// DefaultClientConfiguration returns the default ClientConfiguration, based on viper's patterns of file -> env -> flags
func DefaultClientConfiguration() *ClientConfiguration {
	return &ClientConfiguration{
		BaseURL:  viper.GetString("jira.base_url"),
		Username: viper.GetString("jira.username"),
		Password: viper.GetString("jira.password"),
		Debug:    viper.GetBool("debug"),
		Insecure: viper.GetBool("tls.insecure"),
	}
}
