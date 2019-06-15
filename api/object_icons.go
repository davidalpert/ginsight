package api

import (
	"fmt"
	"strings"
)

type ObjectIcon struct {
	ID        int    `json:"id"`                  // The id
	Name      string `json:"name"`                // The name
	Url16     string `json:"url16"`               // The URL to icon (16x16)
	Url48     string `json:"url48"`               // The URL to icon (48x48)
	Removable bool   `json:"removable,omitempty"` // If removable
}

func (c *Client) GetIconByID(id string) (*ObjectIcon, error) {
	full_url := fmt.Sprintf(c.BaseURL+"/rest/insight/1.0/icon/%s", id)
	var result ObjectIcon
	response, err := c.R().SetResult(&result).Get(full_url)
	if err != nil {
		return nil, err
	}
	if err = validateResponseCodeExact(response, 200); err != nil {
		return nil, err
	}
	return response.Result().(*ObjectIcon), nil
}

func (c *Client) GetGlobalIconByName(name string) (*ObjectIcon, error) {
	icons, err := c.GetGlobalIcons()
	if err != nil {
		return nil, err
	}
	return findIconByName(name, "", icons)
}

func (c *Client) GetSchemaIconByName(schemaKey string, name string) (*ObjectIcon, error) {
	icons, err := c.GetSchemaIcons(schemaKey)
	if err != nil {
		return nil, err
	}
	return findIconByName(name, schemaKey, icons)
}

func findIconByName(name string, scope string, icons *[]ObjectIcon) (*ObjectIcon, error) {
	suggestions := []string{} // just in case
	var foundIcon *ObjectIcon
	for _, icon := range *icons {
		if strings.EqualFold(icon.Name, name) {
			foundIcon = &icon
			break
		} else {
			suggestions = append(suggestions, icon.Name)
		}
	}

	if foundIcon == nil {
		return nil, ObjectIconNotFoundError{
			SchemaIdentifier: scope,
			SearchTerm:       name,
			Suggestions:      &suggestions,
		}
	}

	return foundIcon, nil
}

func (c *Client) GetGlobalIcons() (*[]ObjectIcon, error) {
	full_url := c.BaseURL + "/rest/insight/1.0/icon/global"
	var result []ObjectIcon
	response, err := c.R().SetResult(&result).Get(full_url)
	if err != nil {
		return nil, err
	}
	if err = validateResponseCodeExact(response, 200); err != nil {
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
	if err = validateResponseCodeExact(response, 200); err != nil {
		return nil, err
	}
	return response.Result().(*[]ObjectIcon), nil
}
