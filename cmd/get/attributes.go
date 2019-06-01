// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package get

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"

	format "github.com/davidalpert/ginsight/format"
	insight "github.com/davidalpert/ginsight/insight"
)

// SchemasCmd represents the get command
var CmdGetAttributes = &cobra.Command{
	Use:   "attributes [objectTypeName|objectTypeID]",
	Short: "Get information about JIRA Insight ObjectTypeAttributes in an ObjectSchema",
	Long: `
Retreives a list of ObjectTypeAttributes from the Insight API.
`,
	Args: cobra.RangeArgs(0, 1),
	RunE: getAttributes,
}

func getAttributes(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return getAttributesForSchema(objectSchemaKey)
	} else {
		objectTypeIdentifier := args[0]
		return getAttributesForObjectType(objectSchemaKey, objectTypeIdentifier)
	}
	return nil
}

func getAttributesForSchema(schemaKey string) error {
	fmt.Printf("Looking up ObjectSchema ID for '%s'\n", objectSchemaKey)
	schema, schemaErr := insight.DefaultClient().GetObjectSchemaByKey(schemaKey)
	if schemaErr != nil {
		return schemaErr
	}

	fmt.Printf("Looking up all attributes in ObjectSchema '%s'\n", objectSchemaKey)
	attributes, err := insight.DefaultClient().GetObjectTypeAttributesForSchemaID(strconv.Itoa(schema.ID))
	if err != nil {
		return err
	}

	format.WriteObjectTypeAttributes("Key", objectSchemaKey, nil, attributes) // pick objectIdentifier from each attribute

	return nil
}

func getAttributesForObjectType(schemaKey string, objectTypeIdentifier string) error {
	// TODO: refactor to insight.DefaultClient().GetObjectTypeByNameFromSchemaKey(..) which throws error if more than one match is found
	var objectType *insight.ObjectType
	if _, atoiErr := strconv.Atoi(objectTypeIdentifier); atoiErr == nil {
		fmt.Printf("Looking up ObjectType ID '%s' in schema '%s'\n", objectTypeIdentifier, objectSchemaKey)
		foundType, objectTypeErr := insight.DefaultClient().GetObjectTypeByID(objectTypeIdentifier)
		if objectTypeErr != nil {
			return objectTypeErr
		}
		objectType = foundType
	} else {
		fmt.Printf("Looking up ObjectType ID for '%s' in schema '%s'\n", objectTypeIdentifier, objectSchemaKey)
		foundTypes, objectTypesErr := insight.DefaultClient().GetObjectTypesByNameFromSchemaKey(objectSchemaKey, objectTypeIdentifier)
		if objectTypesErr != nil {
			return objectTypesErr
		}
		if len(*foundTypes) > 1 {
			return &insight.MultipleObjectTypesFoundError{
				SchemaID:       objectSchemaKey,
				ObjectTypeName: objectTypeIdentifier,
				FoundTypes:     foundTypes,
			}
		}
		objectType = &(*foundTypes)[0]
	}
	// END EXTRACT METHODj

	fmt.Printf("Looking up attributes for ObjectType '%s' (%d) in ObjectSchema '%s' ...\n", objectType.Name, objectType.ID, objectSchemaKey)
	attributes, err := insight.DefaultClient().GetObjectTypeAttributesForObjectTypeID(strconv.Itoa(objectType.ID))
	if err != nil {
		return err
	}

	format.WriteObjectTypeAttributes("Key", objectSchemaKey, objectType, attributes) // pick objectIdentifier from each attribute

	return nil
}
