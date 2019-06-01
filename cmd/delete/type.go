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

package delete

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"

	insight "github.com/davidalpert/ginsight/insight"
	//insightFormat "github.com/davidalpert/insight/format"
)

var objectTypeName string
var objectTypeID int

var CmdDeleteType = &cobra.Command{
	Use:   "type",
	Short: "Delete an ObjectType in a JIRA Insight ObjectSchema",
	Long: `
Deletes an existing ObjectType in a JIRA Insight ObjectSchema
`,

	Args: cobra.ExactArgs(1),
	Example: `
  # Delete an ObjectType by id
  insight delete type 45
  
  # Delete an ObjectType by id, bypassing confirmation
  insight delete type 45 --force
  
  # Delete an ObjectType by name in a given schema
  insight delete type Vendor --key IT
  
`,
	RunE: deleteObjectType,
}

func deleteObjectType(cmd *cobra.Command, args []string) error {
	typeIdentifier := args[0] // guaranteed by 'Args: cobra.ExactArgs(1)'

	if _, err := strconv.Atoi(typeIdentifier); err != nil {
		return deleteObjectTypeByNameInSchemaKey(objectSchemaKey, typeIdentifier)
	}

	return deleteObjectTypeByIDInSchemaKey(objectSchemaKey, typeIdentifier)
}

func deleteObjectTypeByNameInSchemaKey(schemaKey string, typeIdentifier string) error {
	if insight.DefaultClient().Debug {
		fmt.Printf("Looking up ObjectTypes by Name '%s' in Schema '%s' ...\n", typeIdentifier, schemaKey)
	}
	objectTypes, err := insight.DefaultClient().GetObjectTypesByNameFromSchemaKey(schemaKey, typeIdentifier)
	if err != nil {
		return err
	}

	if len(*objectTypes) <= 0 {
		return insight.ObjectTypeNotFoundError{
			SearchTerm:       typeIdentifier,
			SchemaIdentifier: schemaKey,
			Suggestions: &[]string{
				"run 'insight get types' to see the full list of types",
			},
		}
	}

	return deleteObjectTypeWithConfirmation(schemaKey, &(*objectTypes)[0])
}

func deleteObjectTypeByIDInSchemaKey(schemaKey string, typeIdentifier string) error {
	if insight.DefaultClient().Debug {
		fmt.Printf("Looking up ObjectType by ID '%s' in Schema '%s' ...\n", typeIdentifier, schemaKey)
	}
	objectType, err := insight.DefaultClient().GetObjectTypeByID(typeIdentifier)
	if err != nil {
		return err
	}

	return deleteObjectTypeWithConfirmation(schemaKey, objectType)
}

func deleteObjectTypeWithConfirmation(schemaIdentifier string, objectType *insight.ObjectType) error {
	fmt.Printf("About to delete ObjectType '%s' from schema '%s' ...\n", objectType.Name, schemaIdentifier)
	fmt.Println("TODO: implement delete")
	return nil
}
