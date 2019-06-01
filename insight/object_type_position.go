package insight

import (
	"fmt"
)

func (c *Client) UpdateObjectTypePosition(objectTypeID string, newPosition int) (*ObjectType, error) {
	body := map[string]interface{}{
		"position": newPosition,
	}
	return c.updateObjectTypePosition(objectTypeID, body)
}

func (c *Client) UpdateObjectTypePositionAndParent(objectTypeID string, newPosition int, newParentTypeID int) (*ObjectType, error) {
	body := map[string]interface{}{
		"position":       newPosition,
		"toObjectTypeId": newParentTypeID,
	}
	return c.updateObjectTypePosition(objectTypeID, body)
}

func (c *Client) updateObjectTypePosition(objectTypeID string, body map[string]interface{}) (*ObjectType, error) {
	url := fmt.Sprintf(c.BaseURL+"/rest/insight/1.0/objecttype/%s/position", objectTypeID)
	response, err := c.R().SetBody(body).SetResult(&ObjectType{}).Put(url)
	if err != nil {
		return nil, err
	}

	if statusErr := validateResponseCodeExact(response, 200); statusErr != nil {
		return nil, statusErr
	}

	return response.Result().(*ObjectType), nil
}
