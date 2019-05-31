package insight

type JiraMe struct {
	Name         string `json:"name"`
	DisplayName  string `json:"displayName`
	EmailAddress string `json:"emailAddress`
	Active       bool   `json:"active`
	TimeZone     string `json:"timeZone`
	Locale       string `json:"locale`
}

// Gets information about the configured Jira credentials.
func (c *Client) GetMe() (*JiraMe, error) {
	// using resty.Request.SetResult allows us to predefine the expected shape of the result
	// provided that the result includes a "Content-Type: application/json", marshalling
	// into our provided struct will be attempted automatically
	response, err := c.R().SetResult(JiraMe{}).Get(c.BaseURL + "/rest/api/2/myself")
	if err != nil {
		return nil, err
	}

	jiraUser := response.Result().(*JiraMe) // casts the Result() interface to a pointer-to-JiraMe
	return jiraUser, nil                    // returns the pointer-to-JiraMe
}
