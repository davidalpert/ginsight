package api

import (
	"fmt"
)

func (c *Client) UpdateObjectTypeParent(objectTypeID string, newParentTypeID *int) (*ObjectType, error) {
	// new parent without a position? let's insert it as the first of it's children
	// newParentTypeID can be null to remove the parent typeID
	return c.UpdateObjectTypeParentAndPosition(objectTypeID, newParentTypeID, 0)
}

func (c *Client) UpdateObjectTypePosition(objectTypeID string, newPosition int) (*ObjectType, error) {
	// only changing the position? we need to keep the same parent object type
	objectType, err := c.GetObjectTypeByID(objectTypeID)
	if err != nil {
		return nil, err
	}

	return c.UpdateObjectTypeParentAndPosition(objectTypeID, &objectType.ParentObjectTypeID, newPosition)
}

func (c *Client) UpdateObjectTypeParentAndPosition(objectTypeID string, newParentTypeID *int, newPosition int) (*ObjectType, error) {
	body := map[string]interface{}{
		"position":       newPosition,
		"toObjectTypeId": newParentTypeID,
	}

	url := fmt.Sprintf(c.BaseURL+"/rest/insight/1.0/objecttype/%s/position", objectTypeID)
	response, err := c.R().SetBody(body).SetResult(&ObjectType{}).Post(url)
	if err != nil {
		return nil, err
	}

	if err := validateResponseCodeExact(response, 200); err != nil {
		return nil, err
	}

	return response.Result().(*ObjectType), nil
}
