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
)

// SchemasCmd represents the get command
var CmdGetFamily = &cobra.Command{
	Use:   "family [identifier]",
	Short: "Get information about an ObjectType family in Jira Insight ObjectSchema",
	Long: `
Retreives a family of ObjectTypes from the Insight API (a parent and it's children).

'identifier' can be either ObjectTypeID or an ObjectType.Name

if 'identifier' is missing return the 'root' family (objects with no parents)
`,
  Example: `
  # get the list of ObjectTypes grouped under the 'Hardware' type family in the IT schema
  ginsight get family hardware --schema IT

  # get the list of ObjectTypes at the root of the IT schema (those with no parent object types) 
  ginsight get family --schema IT
`,
	Args: cobra.MaximumNArgs(1),
	RunE: getTypeFamilyE,
}

func getTypeFamilyE(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		fmt.Printf("Getting the root type family (types with no parents) in '%s'\n", objectSchemaKey)
		return getTypeFamilyByIDInSchemaKey(objectSchemaKey, nil)
	}

	typeIdentifier := args[0]
	if typeID, err := strconv.Atoi(typeIdentifier); err == nil {
		fmt.Printf("Getting the type family for '%s' in '%s'\n", typeIdentifier, objectSchemaKey)
		return getTypeFamilyByIDInSchemaKey(objectSchemaKey, &typeID)
	}

	fmt.Printf("Getting the type family for '%s' in '%s'\n", typeIdentifier, objectSchemaKey)
	return getTypeFamilyByNameInSchemaKey(objectSchemaKey, typeIdentifier)
}

func getTypeFamilyByIDInSchemaKey(objectSchemaKey string, typeID *int) error {
	f, err := api.DefaultClient().GetObjectTypeFamilyForSchemaKey(objectSchemaKey, typeID)
	if err != nil {
		return err
	}

	if f.Parent == nil {
		fmt.Printf("The root family for '%s'\n", f.Schema.Key)
	} else {
		fmt.Printf("The '%s' (%d) family for '%s'\n", f.Parent.Name, f.Parent.ID, f.Schema.Key)
	}

	for _, childType := range f.Children {
		fmt.Printf(" - [%d] %s (%d)\n", childType.Position, childType.Name, childType.ID)
	}

	return nil
}

func getTypeFamilyByNameInSchemaKey(objectSchemaKey string, typeName string) error {
	if typeName == "root" || typeName == "none" || typeName == "nil" {
		return getTypeFamilyByIDInSchemaKey(objectSchemaKey, nil)
	}

	objectType, err := api.DefaultClient().GetObjectTypeByNameFromSchemaKey(objectSchemaKey, typeName)
	if err != nil {
		return err
	}

	return getTypeFamilyByIDInSchemaKey(objectSchemaKey, &objectType.ID)
}
