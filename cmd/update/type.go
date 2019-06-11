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

package create

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"

	api "github.com/davidalpert/ginsight/api"
	insightFormat "github.com/davidalpert/ginsight/format"
)

var objectTypeName string
var objectTypeDescription string
var objectTypeParentObjectTypeID int
var objectTypeParentObjectTypeName string
var objectTypeIconID int

var CmdUpdateType = &cobra.Command{
	Use:   "type",
	Short: "Create an ObjectType in a JIRA Insight ObjectSchema",
	Long: `
Creates a new ObjectType in a JIRA Insight ObjectSchema
`,

	Args: cobra.ExactArgs(1),
	Example: `
  # update a new ObjectType in the IT schema
  insight update type Vendor --schema IT --description "A vendor provides products or services"

  # update a new ObjectType in the IT schema with debug logs enabled
  insight update type Vendor --schema IT --description "A vendor provides products or services" --debug

  # create a new ObjectType in the IT schema with debug logs enabled
  insight update type Vendor --schema IT --description "A vendor provides products or services" --debug

  # update a new ObjectType in the IT schema with a parent ObjectType by parent type Id
  insight update type Vendor --schema IT --description "A vendor provides products or services" --debug

  # update a new ObjectType in the IT schema with debug logs enabled
  insight update type Vendor --schema IT --description "A vendor provides products or services" --debug
`,
	RunE: updateType,
}

func init() {
	CmdUpdateType.Flags().StringVar(&objectTypeName, "new-name", "", "allows updating the name")
	CmdUpdateType.Flags().StringVarP(&objectTypeDescription, "description", "l", "", "description of the object type to create")
	CmdUpdateType.Flags().StringVarP(&objectTypeParentObjectTypeName, "parent-type-name", "p", "", "name of the parent object type")
	CmdUpdateType.Flags().IntVarP(&objectTypeParentObjectTypeID, "parent-type-id", "P", -1, "id of the parent object type")
	CmdUpdateType.Flags().IntVarP(&objectTypeIconID, "icon-id", "i", 1, "id of the icon") // icon is required, 1-indexed, so 1 is the default iconId
	//CmdUpdateType.Flags().MarkHidden("new-name")
}

func updateType(cmd *cobra.Command, args []string) error {
	typeIdentifier := args[0] // guaranteed by 'Args: cobra.ExactArgs(1)'

	typeUpdate := api.ObjectTypeUpdateRequest{
		Name:        typeIdentifier,
		Description: objectTypeDescription,
		IconID:      objectTypeIconID,
	}

	if cmd.Flags().Changed("new-name") {
		typeUpdate.Name = objectTypeName // use the new name
	}

	if _, err := strconv.Atoi(typeIdentifier); err != nil {
		return updateObjectTypeByNameInSchemaKey(objectSchemaKey, typeIdentifier, &typeUpdate)
	}

	return updateObjectTypeByIdInSchemaKey(objectSchemaKey, typeIdentifier, &typeUpdate)
}

func updateObjectTypeByNameInSchemaKey(schemaKey string, typeIdentifier string, update *api.ObjectTypeUpdateRequest) error {
	if api.DefaultClient().Debug {
		fmt.Printf("Looking up ObjectTypes by Name '%s' in Schema '%s' ...\n", typeIdentifier, schemaKey)
	}

	foundTypes, err := api.DefaultClient().GetObjectTypesByNameFromSchemaKey(schemaKey, typeIdentifier)
	if err != nil {
		return err
	}

	if len(*foundTypes) < 1 {
		return api.ObjectTypeNotFoundError{
			SearchTerm:       typeIdentifier,
			SchemaIdentifier: schemaKey,
			Suggestions: &[]string{
				"run 'insight get types' to see the full list of types",
			},
		}
	}

	if len(*foundTypes) > 1 {
		return &api.MultipleObjectTypesFoundError{
			SchemaID:       schemaKey,
			ObjectTypeName: typeIdentifier,
			FoundTypes:     foundTypes,
		}
	}

	objectType := (*foundTypes)[0]

	return applyUpdates(schemaKey, &objectType, update)
}

func updateObjectTypeByIdInSchemaKey(schemaKey string, typeIdentifier string, update *api.ObjectTypeUpdateRequest) error {
	if api.DefaultClient().Debug {
		fmt.Printf("Looking up ObjectType by Id '%s' in Schema '%s' ...\n", typeIdentifier, schemaKey)
	}

	objectType, err := api.DefaultClient().GetObjectTypeByID(typeIdentifier)
	if err != nil {
		return err
	}

	// don't change the name when we looked it up by ID and haven't set the name to something new
	if update.Name == typeIdentifier {
		update.Name = objectType.Name
	}

	return applyUpdates(schemaKey, objectType, update)
}

func applyUpdates(schemaKey string, objectType *api.ObjectType, update *api.ObjectTypeUpdateRequest) error {
	updatedObjectType, err := api.DefaultClient().UpdateObjectType(strconv.Itoa(objectType.ID), update)
	if err != nil {
		return err
	}

	fmt.Printf("\nUpdated ObjectType %s (%d) in ObjectSchema %s\n\n", updatedObjectType.Name, updatedObjectType.ID, objectSchemaKey)

	insightFormat.WriteObjectType("Key", objectSchemaKey, updatedObjectType)

	return nil
}
