package api

import (
	"fmt"
	"strings"
)

var AttributeTypeIDsByName = map[string]int{
	"default":    0,
	"object":     1,
	"user":       2,
	"confluence": 3,
	"group":      4,
	"version":    5,
	"project":    6,
	"status":     7,
}

func AttributeTypeNameToID(name string) int {
	nameLower := strings.ToLower(name)
	for typeName, id := range AttributeTypeIDsByName {
		if strings.EqualFold(nameLower, typeName) {
			return id
		}
	}
	return -1
}

// The DefaultType if type is set to "Default"
var AttributeDefaultTypeIDsByName = map[string]int{
	"text":      0,
	"integer":   1,
	"boolean":   2,
	"double":    3,
	"date":      4,
	"time":      5,
	"date_time": 6,
	"url":       7,
	"email":     8,
	"textarea":  9,
	"select":    10,
}

func AttributeDefaultTypeIDToName(id int) string {
	for typeName, typeID := range AttributeDefaultTypeIDsByName {
		if typeID == id {
			return typeName
		}
	}
	return fmt.Sprintf("unknown (%d)", id)
}

type ObjectTypeAttribute struct {
	ID          int         `json:"id"`
	ObjectType  *ObjectType `json:"objectType,omitempty"`
	Name        string      `json:"name"`
	Label       bool        `json:"label"`
	TypeID      int         `json:"type"`
	DefaultType struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"defaultType"`
	Editable                bool   `json:"editable"`
	System                  bool   `json:"system"`
	Sortable                bool   `json:"sortable"`
	Summable                bool   `json:"summable"`
	MinimumCardinality      int    `json:"minimumCardinality"`
	MaximumCardinality      int    `json:"maximumCardinality"`
	Removable               bool   `json:"removable"`
	Hidden                  bool   `json:"hidden"`
	IncludeChildObjectTypes bool   `json:"includeChildObjectTypes"`
	UniqueAttribute         bool   `json:"uniqueAttribute"`
	Options                 string `json:"options"`
	Position                int    `json:"position"`
	Description             string `json:"description,omitempty"`
	AdditionalValue         string `json:"additionalValue,omitempty"`
}

type ObjectTypeLabelAttributeUpdateRequest struct {
	Name          string `json:"name"`
	Description   string `json:"description,omitempty"`
}

type ObjectTypeDefaultAttributeCreateRequest struct {
	Name          string `json:"name"`
	Description   string `json:"description,omitempty"`
	TypeID        int    `json:"type"`
	DefaultTypeID int    `json:"defaultTypeId"`

	AdditionalValue string `json:"additionalValue,omitempty"`
	Options         string `json:"options,omitempty"`
}

type ObjectTypeObjectAttributeCreateRequest struct {
	Name         string `json:"name"`
	TypeID       int    `json:"type"`
	ObjectTypeID int    `json:"typeValue"` // ID of the referenced object type

	ReferenceTypeID int `json:"additionalValue,omitempty"` // ID of the reference type
}

type ObjectTypeUserAttributeCreateRequest struct {
	Name            string `json:"name"`
	TypeID          int    `json:"type"`
	AdditionalValue string `json:"additionalValue"` // SHOW_PROFILE, HIDE_PROFILE

	JiraGroups []string `json:"typeValueMulti,omitempty"`
}

func (c *Client) GetObjectTypeAttributesForSchemaID(objectSchemaID string) (*[]ObjectTypeAttribute, error) {
	var result []ObjectTypeAttribute
	response, err := c.R().SetResult(&result).Get(fmt.Sprintf(c.BaseURL+"/rest/insight/1.0/objectschema/%s/attributes", objectSchemaID))
	if err != nil {
		return nil, err
	}
	return response.Result().(*[]ObjectTypeAttribute), nil
}

func (c *Client) GetObjectTypeAttributesForObjectTypeID(objectTypeID string) (*[]ObjectTypeAttribute, error) {
	var result []ObjectTypeAttribute
	response, err := c.R().SetResult(&result).Get(fmt.Sprintf(c.BaseURL+"/rest/insight/1.0/objecttype/%s/attributes", objectTypeID))
	if err != nil {
		return nil, err
	}
	return response.Result().(*[]ObjectTypeAttribute), nil
}

func (c *Client) UpdateLabelAttributeForObjectTypeID(objectTypeID string, name string, description string) (*ObjectTypeAttribute, error) {
	var result ObjectTypeAttribute
	body := &ObjectTypeLabelAttributeUpdateRequest{
		Name: name,
		Description: description,
	}

	labelAttribute, err := c.GetLabelAttributeForObjectTypeID(objectTypeID)
	if err != nil {
		return nil, err
	}

	response, err := c.R().SetBody(body).SetResult(&result).Put(fmt.Sprintf(c.BaseURL+"/rest/insight/1.0/objecttypeattribute/%s/%d", objectTypeID, labelAttribute.ID))
	if err != nil {
		return nil, err
	}

	if err = validateResponseCodeExact(response, 200); err != nil {
		return nil, err
	}

	return response.Result().(*ObjectTypeAttribute), nil
}

func (c *Client) GetLabelAttributeForObjectTypeID(objectTypeID string) (*ObjectTypeAttribute, error) {
	attributes, err := c.GetObjectTypeAttributesForObjectTypeID(objectTypeID)
	if err != nil {
		return nil, err
	}

	for _, a := range *attributes {
		if a.Editable && a.Label {
			return &a, nil
		}
	}

	return nil, fmt.Errorf("Could not find the Label attribute for ObjectTypeID [%s]!", objectTypeID)
}

func (c *Client) GetEditableObjectTypeAttributesForObjectTypeID(objectTypeID string) (*[]ObjectTypeAttribute, error) {
	attributes, err := c.GetObjectTypeAttributesForObjectTypeID(objectTypeID)
	if err != nil {
		return nil, err
	}

	editableAttributes := make([]ObjectTypeAttribute, 0)
	for _, a := range *attributes {
		if a.Editable {
			editableAttributes = append(editableAttributes, a)
		}
	}

	return &editableAttributes, nil
}

func (c *Client) CreateObjectTypeDefaultAttribute(objectTypeID string, body *ObjectTypeDefaultAttributeCreateRequest) (*ObjectTypeAttribute, error) {
	response, err := c.R().SetBody(body).SetResult(&ObjectTypeAttribute{}).
		Post(fmt.Sprintf(c.BaseURL+"/rest/insight/1.0/objecttypeattribute/%s", objectTypeID))
	if err != nil {
		return nil, err
	}
	return response.Result().(*ObjectTypeAttribute), nil
}

func (c *Client) CreateObjectTypeObjectAttribute(objectTypeID string, body *ObjectTypeDefaultAttributeCreateRequest) (*ObjectTypeAttribute, error) {
	response, err := c.R().SetBody(body).SetResult(&ObjectTypeAttribute{}).
		Post(fmt.Sprintf(c.BaseURL+"/rest/insight/1.0/objecttypeattribute/%s", objectTypeID))
	if err != nil {
		return nil, err
	}
	return response.Result().(*ObjectTypeAttribute), nil
}
