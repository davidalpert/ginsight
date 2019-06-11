package api

// JiraMe models the response of Jira's /rest/api/2/myself endpoint
type JiraMe struct {
	Name         string `json:"name"`
	DisplayName  string `json:"displayName`
	EmailAddress string `json:"emailAddress`
	Active       bool   `json:"active`
	TimeZone     string `json:"timeZone`
	Locale       string `json:"locale`
}

// GetMe fetches information about the configured Jira credentials.
func (c *Client) GetMe() (*JiraMe, error) {
	response, err := c.R().SetResult(JiraMe{}).Get(c.BaseURL + "/rest/api/2/myself")
	if err != nil {
		return nil, err
	}
	err = validateResponseCodeInRange(response, 200, 300)
	if err != nil {
		return nil, err
	}

	return response.Result().(*JiraMe), err
}
