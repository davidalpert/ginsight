package api

import (
	"fmt"

	resty "gopkg.in/resty.v1"
)

// -----------------------------------------------

// ClientError wraps a restyResponse
type ClientError struct {
	Response *resty.Response
}

// Error provides the error response as a string
func (e *ClientError) Error() string {
	if e.Response.StatusCode() == 403 {
		return string("403 Unauthorized")
	}
	return fmt.Sprintf("%d %s\n", e.Response.StatusCode(), e.Response.String())
}

// -----------------------------------------------

// ObjectSchemaNotFoundError represents a schema not found
type ObjectSchemaNotFoundError struct {
	SearchTerm  string
	Suggestions []string
}

// Error provides the error response as a string
func (e ObjectSchemaNotFoundError) Error() string {
	return fmt.Sprintf("Did not find schema '%s'\n\nAre you looking for one of these schemas? %s\n", e.SearchTerm, e.Suggestions)
}

// -----------------------------------------------

// ObjectSchemaKeyMismatchError represents an invalid attempt to update a schema key
type ObjectSchemaKeyMismatchError struct {
	SchemaId string
	ExistingKey  string
	NewKey string
}

// Error provides the error response as a string
func (e ObjectSchemaKeyMismatchError) Error() string {
	return fmt.Sprintf("Invalid attempt to update schema key for Schema #%s from '%s' to '%s'\n\nthe insight API does not allow changing a schema key after it has been created.\n", e.SchemaId, e.ExistingKey, e.NewKey)
}

// -----------------------------------------------

// ObjectTypeNotFoundError represents an ObjectType not found
type ObjectTypeNotFoundError struct {
	SearchTerm       string
	SchemaIdentifier string
	Suggestions      *[]string
}

func (e ObjectTypeNotFoundError) Error() string {
	return fmt.Sprintf("Did not find type '%s' in schema '%s'\n\nAre you looking for one of these types? %s\n", e.SearchTerm, e.SchemaIdentifier, e.Suggestions)
}

// -----------------------------------------------

// MultipleObjectTypesFoundError represents multiple types found when we expected only a single type
type MultipleObjectTypesFoundError struct {
	SchemaID       string
	ObjectTypeName string
	FoundTypes     *[]ObjectType
}

func (e *MultipleObjectTypesFoundError) Error() string {
	var suggestions []string
	for _, objectType := range *e.FoundTypes {
		suggestions = append(suggestions, fmt.Sprintf("- %d '%s'\n", objectType.ID, objectType.Name))
	}
	return fmt.Sprintf(`Found more than one ObjectType matching the criteria; there are %d ObjectTypes with SchemaId %s named %s.

%s

You may need to try again using the Id of the ObjectType you want.

`, len(*e.FoundTypes), e.SchemaID, e.ObjectTypeName, suggestions)
}

// -----------------------------------------------

// ObjectIconNotFoundError represents an ObjectIcon not found
type ObjectIconNotFoundError struct {
	SearchTerm       string
	SchemaIdentifier string
	Suggestions      *[]string
}

func (e ObjectIconNotFoundError) Error() string {
	if e.SchemaIdentifier == "" {
		return fmt.Sprintf("Did not find icon '%s' in the global icon set.\n\nAre you looking for one of these types? %s\n", e.SearchTerm, e.Suggestions)
	} else {
		return fmt.Sprintf("Did not find icon '%s' in ObjectSchema '%s'\n\nAre you looking for one of these types? %s\n", e.SearchTerm, e.SchemaIdentifier, e.Suggestions)
	}
}
