package api

import (
	"fmt"
	"sort"
)

type ObjectTypeFamily struct {
	Schema   *ObjectSchema
	Parent   *ObjectType // null for the 'root' family
	Children []ObjectType
}

func (c *Client) GetObjectTypeFamilyForSchemaKey(schemaKey string, parentTypeID *int) (*ObjectTypeFamily, error) {
	objectSchema, err := c.GetObjectSchemaByKey(schemaKey)
	if err != nil {
		return nil, err
	}

	family := &ObjectTypeFamily{
		Schema: objectSchema,
	}

	allObjectTypes, err := c.GetObjectTypesForSchemaID(objectSchema.IDString())
	if err != nil {
		return nil, err
	}

	if c.Debug {
		fmt.Println("Found object types:")
		for _, t := range *allObjectTypes {
			if t.ParentObjectTypeID == nil {
				fmt.Printf("- '%s' (%d) - position (%d) - parent: '%v'\n", t.Name, t.ID, t.Position, t.ParentObjectTypeID)
			} else {
				fmt.Printf("- '%s' (%d) - position (%d) - parent: '%d'\n", t.Name, t.ID, t.Position, *t.ParentObjectTypeID)
			}
		}
	}

	childTypes := []ObjectType{}
	for _, t := range *allObjectTypes {
		if parentTypeID != nil && t.ID == *parentTypeID {
			// found the family parent
			family.Parent = &t
		}
		if (parentTypeID == nil && t.ParentObjectTypeID == nil) || (parentTypeID != nil && t.ParentObjectTypeID != nil && *parentTypeID == *t.ParentObjectTypeID) {
			// found one of this family's children
			childTypes = append(childTypes, t)
		}
	}

	if c.Debug {
		fmt.Println("Found child types:")
		for _, t := range childTypes {
			if t.ParentObjectTypeID == nil {
				fmt.Printf("- '%s' (%d) - position (%d) - parent: '%v'\n", t.Name, t.ID, t.Position, t.ParentObjectTypeID)
			} else {
				fmt.Printf("- '%s' (%d) - position (%d) - parent: '%d'\n", t.Name, t.ID, t.Position, *t.ParentObjectTypeID)
			}
		}
	}

	sort.Sort(ByObjectTypePosition(childTypes))

	if c.Debug {
		fmt.Println("child types sorted by position:")
		for _, t := range childTypes {
			if t.ParentObjectTypeID == nil {
				fmt.Printf("- '%s' (%d) - position (%d) - parent: '%v'\n", t.Name, t.ID, t.Position, t.ParentObjectTypeID)
			} else {
				fmt.Printf("- '%s' (%d) - position (%d) - parent: '%d'\n", t.Name, t.ID, t.Position, *t.ParentObjectTypeID)
			}
		}
	}

	for _, t := range childTypes {
		family.Children = append(family.Children, t)
		//family.Children[t.Position] = *t
	}

	return family, nil
}

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

	return c.UpdateObjectTypeParentAndPosition(objectTypeID, objectType.ParentObjectTypeID, newPosition)
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
