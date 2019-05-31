package insight

import (
	"fmt"

	"gopkg.in/resty.v1"
)

// -----------------------------------------------

// Represents an Insight API Client error
type ClientError struct {
	Response *resty.Response
}

// Implements the error interface
func (e *ClientError) Error() string {
	return string(e.Response.String())
}

// -----------------------------------------------

type ObjectSchemaNotFoundError struct {
	SearchTerm  string
	Suggestions []string
}

func (e ObjectSchemaNotFoundError) Error() string {
	return fmt.Sprintf("Did not find schema '%s'\n\nAre you looking for one of these schemas? %s\n", e.SearchTerm, e.Suggestions)
}

// -----------------------------------------------

type ObjectTypeNotFoundError struct {
	SearchTerm       string
	SchemaIdentifier string
	Suggestions      *[]string
}

func (e ObjectTypeNotFoundError) Error() string {
	return fmt.Sprintf("Did not find type '%s' in schema '%s'\n\nAre you looking for one of these types? %s\n", e.SearchTerm, e.SchemaIdentifier, e.Suggestions)
}

// -----------------------------------------------

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
