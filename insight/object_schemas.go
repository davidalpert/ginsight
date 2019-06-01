package insight

import (
	"fmt"
	"log"
	"strings"
)

type ObjectSchema struct {
	ID              int    `json:"id"`              // The id
	Name            string `json:"name"`            // The name
	Key             string `json:"objectSchemaKey"` // The key
	Description     string `json:"description"`     // The description
	Status          string `json:"status"`          // The status
	Created         string `json:"created"`         // The created date
	Updated         string `json:"updated"`         // The updated date
	ObjectTypeCount int    `json:"objectTypeCount"` // Number of object types
	ObjectCount     int    `json:"objectCount"`     // Number of objects
}

type ObjectSchemaList struct {
	Schemas []ObjectSchema `json:"objectschemas"`
}

type ObjectSchemaCreateUpdateRequest struct {
	Name        string `json:"name"`            // The name
	Key         string `json:"objectSchemaKey"` // The key
	Description string `json:"description"`     // The description
}

// Get All Schemas
func (c *Client) GetObjectSchemas() (*ObjectSchemaList, error) {
	//var sampleResponse []ObjectSchema
	response, err := c.R().SetResult(ObjectSchemaList{}).Get(c.BaseURL + "/rest/insight/1.0/objectschema/list")
	if err != nil {
		return nil, err
	}

	if statusErr := validateResponseCodeExact(response, 200); statusErr != nil {
		return nil, statusErr
	}

	schemaList := response.Result().(*ObjectSchemaList)

	if c.Debug {
		log.Printf("\nschemaList: %T (%d schemas)\n", *schemaList, len(schemaList.Schemas))
	}

	return schemaList, nil
}

// Get One Schema by ID
func (c *Client) GetObjectSchemaById(id string) (*ObjectSchema, error) {
	response, err := c.R().SetResult(&ObjectSchema{}).Get(c.BaseURL + "/rest/insight/1.0/objectschema/" + id)
	if err != nil {
		return nil, err
	}

	if statusErr := validateResponseCodeExact(response, 200); statusErr != nil {
		return nil, statusErr
	}

	return response.Result().(*ObjectSchema), nil
}

// Create an ObjectSchema
func (c *Client) CreateSchema(body *ObjectSchemaCreateUpdateRequest) (*ObjectSchema, error) {
	response, err := c.R().SetBody(body).SetResult(&ObjectSchema{}).Post(c.BaseURL + "/rest/insight/1.0/objectschema/create")
	if err != nil {
		return nil, err
	}

	if statusErr := validateResponseCodeExact(response, 201); statusErr != nil {
		return nil, statusErr
	}

	return response.Result().(*ObjectSchema), nil
}

// Update an ObjectSchema
func (c *Client) UpdateSchema(objectSchemaId string, body *ObjectSchemaCreateUpdateRequest) (*ObjectSchema, error) {
	response, err := c.R().SetBody(body).SetResult(&ObjectSchema{}).Put(c.BaseURL + "/rest/insight/1.0/objectschema/" + objectSchemaId)
	if err != nil {
		return nil, err
	}

	if statusErr := validateResponseCodeExact(response, 200); statusErr != nil {
		return nil, statusErr
	}

	return response.Result().(*ObjectSchema), nil
}

// Delete an ObjectSchema
func (c *Client) DeleteSchema(objectSchemaId string) error {
  response, err := c.R().Delete(c.BaseURL + "/rest/insight/1.0/objectschema/" + objectSchemaId)

	if statusErr := validateResponseCodeExact(response, 200); statusErr != nil {
		return statusErr
  }

	return err
}

// Helpers ---------------------------------

// Get Schema by KEY
func (c *Client) GetObjectSchemaByKey(key string) (*ObjectSchema, error) {
	response, objectSchemasErr := c.GetObjectSchemas()
	if objectSchemasErr != nil {
		return nil, objectSchemasErr
	}

	suggestions := []string{} // just in case
	var foundSchema *ObjectSchema
	for _, schema := range response.Schemas {
		if strings.EqualFold(schema.Key, key) {
			foundSchema = &schema
			break
		} else {
			suggestions = append(suggestions, schema.Key)
		}
	}

	if foundSchema == nil {
		return nil, &ObjectSchemaNotFoundError{SearchTerm: key, Suggestions: suggestions}
	}

	return foundSchema, nil
}

// Delete Schema By KEY
func (c *Client) DeleteSchemaByKey(key string) error {
	schema, err := c.GetObjectSchemaByKey(key)
	if err != nil {
		return err
	}

	return c.DeleteSchema(fmt.Sprintf("%d", schema.ID))
}
