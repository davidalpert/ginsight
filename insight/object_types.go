package insight

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type ObjectType struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Type        int    `json:"type"`
	Description string `json:"description"`
	Icon        struct {
		ID    int    `json:"id"`
		Name  string `json:"name"`
		URL16 string `json:"url16"`
		URL48 string `json:"url48"`
	} `json:"icon"`
	Position                  int    `json:"position"`
	Created                   string `json:"created"`
	Updated                   string `json:"updated"`
	ObjectCount               int    `json:"objectCount"`
	ParentObjectTypeID        int    `json:"parentObjectTypeId"`
	ObjectSchemaID            int    `json:"objectSchemaId"`
	Inherited                 bool   `json:"inherited"`
	AbstractObjectType        bool   `json:"abstractObjectType"`
	ParentObjectTypeInherited bool   `json:"parentObjectTypeInherited"`
}

type ObjectTypeCreateRequest struct {
	Name               string `json:"name"`                         // The name
	Description        string `json:"description"`                  // The description
	IconID             int    `json:"iconId"`                       // The icon id
	ParentObjectTypeID int    `json:"parentObjectTypeId,omitempty"` // The parent object type id
	ObjectSchemaID     int    `json:"objectSchemaId"`               // The Object Schema id
}

type ObjectTypeUpdateRequest struct {
	Name        string `json:"name"`        // The name
	Description string `json:"description"` // The description
	IconID      int    `json:"iconId"`      // The icon id
}

// Get all ObjectTypes
func (c *Client) GetObjectTypesForSchemaIDInt(schemaID int) (*[]ObjectType, error) {
	return c.GetObjectTypesForSchemaID(strconv.Itoa(schemaID))
}

// Get all ObjectTypes
func (c *Client) GetObjectTypesForSchemaID(schemaID string) (*[]ObjectType, error) {
	if c.Debug {
		log.Println("GetObjectTypesForSchemaID")
	}

	full_url := fmt.Sprintf(c.BaseURL+"/rest/insight/1.0/objectschema/%s/objecttypes/flat", schemaID)
	var result []ObjectType
	response, err := c.R().SetResult(&result).Get(full_url)
	if err != nil {
		return nil, err
	}

	if statusErr := validateResponseCodeExact(response, 200); statusErr != nil {
		return nil, statusErr
	}

	return response.Result().(*[]ObjectType), nil
}

// Get all ObjectTypes for Schema
func (c *Client) GetObjectTypesForSchemaKey(key string) (*[]ObjectType, error) {
	if c.Debug {
		log.Printf("GetObjectTypesForSchemaKey %s\n", key)
	}

	schema, err := c.GetObjectSchemaByKey(key)
	if err != nil {
		return nil, err
	}

	return c.GetObjectTypesForSchemaIDInt(schema.ID)
}

func (c *Client) GetObjectTypeByID(id string) (*ObjectType, error) {
	if c.Debug {
		log.Println("GetObjectTypeByID")
  }

	response, err := c.R().SetResult(&ObjectType{}).Get(c.BaseURL + "/rest/insight/1.0/objecttype/" + id)
	if err != nil {
		return nil, err
  }
  
  if statusErr := validateResponseCodeExact(response, 200); statusErr != nil {
		return statusErr
  }

	return response.Result().(*ObjectType), nil
}

func (c *Client) GetObjectTypesByNameFromSchemaIDInt(schemaID int, name string) (*[]ObjectType, error) {
	if c.Debug {
		log.Println("GetObjectTypesByNameFromSchemaIDInt")
	}
	return c.GetObjectTypesByNameFromSchemaID(strconv.Itoa(schemaID), name)
}

func (c *Client) GetObjectTypesByNameFromSchemaID(schemaID string, name string) (*[]ObjectType, error) {
	if c.Debug {
		log.Printf("GetObjectTypesByNameFromSchemaID: %s %s", schemaID, name)
	}
	objectTypes, err := c.GetObjectTypesForSchemaID(schemaID)
	if err != nil {
		return nil, err
	}
	return c.lookupObjectTypesByName(schemaID, objectTypes, name)
}

func (c *Client) GetObjectTypesByNameFromSchemaKey(key string, typeName string) (*[]ObjectType, error) {
	if c.Debug {
		log.Printf("GetObjectTypesByNameFromSchemaKey: schema '%s' typeName '%s'\n", key, typeName)
	}
	objectTypes, err := c.GetObjectTypesForSchemaKey(key)
	if err != nil {
		return nil, err
	}
	return c.lookupObjectTypesByName(key, objectTypes, typeName)
}

func (c *Client) lookupObjectTypesByName(schemaIdentifier string, objectTypes *[]ObjectType, name string) (*[]ObjectType, error) {
	if c.Debug {
		log.Printf("Searching %d types for %s in schema %s", len(*objectTypes), name, schemaIdentifier)
	}
	var foundTypes []ObjectType
	var suggestions []string // just in case
	for _, objectType := range *objectTypes {
		if strings.EqualFold(objectType.Name, name) {
			foundTypes = append(foundTypes, objectType)
		} else {
			suggestions = append(suggestions, objectType.Name)
		}
	}

	if len(foundTypes) == 0 {
		log.Println("404 - found no matching types")
		return nil, &ObjectTypeNotFoundError{
			SchemaIdentifier: schemaIdentifier,
			SearchTerm:       name,
			Suggestions:      &suggestions,
		}
	}

	return &foundTypes, nil
}

func (c *Client) CreateObjectType(body *ObjectTypeCreateRequest) (ObjectType, error) {
	response, err := c.R().SetBody(body).SetResult(&ObjectType{}).Post(c.BaseURL + "/rest/insight/1.0/objecttype/create")
	if err != nil {
		return ObjectType{}, err
  }

  if statusErr := validateResponseCodeExact(response, 201); statusErr != nil {
		return statusErr
  }

	return *(response.Result().(*ObjectType)), nil
}

func (c *Client) UpdateObjectType(objectTypeID string, body *ObjectTypeUpdateRequest) (ObjectType, error) {
	response, err := c.R().SetBody(body).SetResult(&ObjectType{}).Put(c.BaseURL + "/rest/insight/1.0/objecttype/" + objectTypeID)
	if err != nil {
		return ObjectType{}, err
  }

  if statusErr := validateResponseCodeExact(response, 201); statusErr != nil {
		return statusErr
  }

	return *(response.Result().(*ObjectType)), nil
}

func (c *Client) DeleteObjectType(objectTypeID string) error {
  response, err := c.R().Delete(c.BaseURL + "/rest/insight/1.0/objecttype/" + objectTypeID)

  if statusErr := validateResponseCodeExact(response, 200); statusErr != nil {
		return statusErr
  }

	return err
}

func (c *Client) DeleteObjectTypeNameInSchemaID(schemaID string, objectTypeName string) error {
	objectTypes, err := c.GetObjectTypesByNameFromSchemaID(schemaID, objectTypeName)
	if err != nil {
		return err
	}

	return c.deleteObjectTypeByNameInSchema(schemaID, objectTypeName, objectTypes)
}

func (c *Client) DeleteObjectTypeByNameInSchemaByKey(key string, objectTypeName string) error {
	objectTypes, err := c.GetObjectTypesByNameFromSchemaKey(key, objectTypeName)
	if err != nil {
		return err
	}

	return c.deleteObjectTypeByNameInSchema(key, objectTypeName, objectTypes)
}

func (c *Client) deleteObjectTypeByNameInSchema(schemaIdentifier string, objectTypeName string, foundTypes *[]ObjectType) error {
	if len(*foundTypes) > 1 {
		return &MultipleObjectTypesFoundError{
			SchemaID:       schemaIdentifier,
			ObjectTypeName: objectTypeName,
			FoundTypes:     foundTypes,
		}
	}

	objectType := (*foundTypes)[0]

	fmt.Printf("Deleting ObjectType %s (%d) ...", objectType.Name, objectType.ID)
	deleteErr := c.DeleteObjectType(fmt.Sprintf("%d", objectType.ID))

	return deleteErr
}
