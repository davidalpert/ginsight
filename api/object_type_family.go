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
	}

	return family, nil
}
