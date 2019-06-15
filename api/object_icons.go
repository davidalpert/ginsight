package api

import (
	"fmt"
)

type ObjectIcon struct {
	ID        int    `json:"id"`                  // The id
	Name      string `json:"name"`                // The name
	Url16     string `json:"url16"`               // The URL to icon (16x16)
	Url48     string `json:"url48"`               // The URL to icon (48x48)
	Removable bool   `json:"removable,omitempty"` // If removable
}

func (c *Client) GetGlobalIcons() (*[]ObjectIcon, error) {
	full_url := c.BaseURL + "/rest/insight/1.0/icon/global"
	var result []ObjectIcon
	response, err := c.R().SetResult(&result).Get(full_url)
	if err != nil {
		return nil, err
	}
	return response.Result().(*[]ObjectIcon), nil
}

func (c *Client) GetSchemaIcons(objectSchemaKey string) (*[]ObjectIcon, error) {
	schema, err := c.GetObjectSchemaByKey(objectSchemaKey)
	if err != nil {
		return nil, err
	}

	full_url := fmt.Sprintf(c.BaseURL+"/rest/insight/1.0/icon/objectschema/%d", schema.ID)
	var result []ObjectIcon
	response, err := c.R().SetResult(&result).Get(full_url)
	if err != nil {
		return nil, err
	}
	return response.Result().(*[]ObjectIcon), nil
}
