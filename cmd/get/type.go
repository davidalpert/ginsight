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

	api "github.com/davidalpert/ginsight/api"
	format "github.com/davidalpert/ginsight/format"
)

// SchemasCmd represents the get command
var CmdGetType = &cobra.Command{
	Use:   "type identifier",
	Short: "Get information about one (or more) JIRA Insight ObjectTypes",
	Long: `
Retreives one (or more; see examples) of ObjectTypes from the Insight API.

'identifier' can be either ObjectTypeID or an ObjectType.Name
`,
	Args: cobra.ExactArgs(1),
	//PreRun: get
	RunE: getTypeBySchemaKey,
}

func init() {}

func getTypeBySchemaKey(cmd *cobra.Command, args []string) error {
	typeIdentifier := args[0]

	if _, err := strconv.Atoi(typeIdentifier); err != nil {
		return getTypesByNameInSchemaKey(objectSchemaKey, typeIdentifier)
	}

	return getTypeByIDInSchemaKey(objectSchemaKey, typeIdentifier)
}

func getTypesByNameInSchemaKey(schemaKey string, typeIdentifier string) error {
	fmt.Printf("Looking up ObjectTypes by Name '%s' in Schema '%s'\n", typeIdentifier, schemaKey)
	objectTypes, err := api.DefaultClient().GetObjectTypesByNameFromSchemaKey(schemaKey, typeIdentifier)
	if err != nil {
		return err
	}

	if len(*objectTypes) <= 0 {
		return api.ObjectTypeNotFoundError{
			SearchTerm:       typeIdentifier,
			SchemaIdentifier: schemaKey,
			Suggestions: &[]string{
				"run 'insight get types' to see the full list of types",
			},
		}
	}

	format.WriteObjectTypes("Key", objectSchemaKey, objectTypes)
	return nil
}

func getTypeByIDInSchemaKey(schemaKey string, typeIdentifier string) error {
	fmt.Printf("Looking up ObjectType by ID '%s' in Schema '%s'\n", typeIdentifier, schemaKey)
	objectType, err := api.DefaultClient().GetObjectTypeByID(typeIdentifier)
	if err != nil {
		return err
	}

	format.WriteObjectType("Key", objectSchemaKey, objectType)
	return nil
}
