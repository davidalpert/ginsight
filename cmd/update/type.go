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
var objectTypeIconName string

var CmdUpdateType = &cobra.Command{
	Use:   "type",
	Short: "Create an ObjectType in a JIRA Insight ObjectSchema",
	Long: `
Creates a new ObjectType in a JIRA Insight ObjectSchema
`,

	Args: cobra.ExactArgs(1),
	Example: `
  # update an ObjectType's name in the IT schema
  insight update type Vendor --schema IT --name ServiceProvider

  # update an ObjectType's description in the IT schema
  insight update type Vendor --schema IT --description "A vendor provides products or services"

  # update an ObjectType's icon by ID in the IT schema
  insight update type Vendor --schema IT --icon-id 42

  # update an ObjectType's icon by name in the IT schema
  insight update type Vendor --schema IT --icon-name building
`,
	RunE: updateType,
}

func init() {
	CmdUpdateType.Flags().StringVar(&objectTypeName, "name", "", "allows updating the name")
	CmdUpdateType.Flags().StringVarP(&objectTypeDescription, "description", "l", "", "description of the object type to create")
	CmdUpdateType.Flags().StringVarP(&objectTypeParentObjectTypeName, "parent-type-name", "p", "", "name of the parent object type")
	CmdUpdateType.Flags().IntVarP(&objectTypeParentObjectTypeID, "parent-type-id", "P", -1, "id of the parent object type")
	CmdUpdateType.Flags().IntVarP(&objectTypeIconID, "icon-id", "i", 1, "id of the icon")     // icon is required, 1-indexed, so 1 is the default iconId
	CmdUpdateType.Flags().StringVar(&objectTypeIconName, "icon-name", "", "name of the icon") // icon will be looked up against the global set
}

func updateType(cmd *cobra.Command, args []string) error {
	typeIdentifier := args[0] // guaranteed by 'Args: cobra.ExactArgs(1)'

	if _, err := strconv.Atoi(typeIdentifier); err != nil {
		return updateObjectTypeByNameInSchemaKey(cmd, objectSchemaKey, typeIdentifier)
	}

	return updateObjectTypeByIdInSchemaKey(cmd, objectSchemaKey, typeIdentifier)
}

func updateObjectTypeByNameInSchemaKey(cmd *cobra.Command, schemaKey string, typeIdentifier string) error {
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

	return applyUpdates(cmd, schemaKey, &objectType)
}

func updateObjectTypeByIdInSchemaKey(cmd *cobra.Command, schemaKey string, typeIdentifier string) error {
	if api.DefaultClient().Debug {
		fmt.Printf("Looking up ObjectType by Id '%s' in Schema '%s' ...\n", typeIdentifier, schemaKey)
	}

	objectType, err := api.DefaultClient().GetObjectTypeByID(typeIdentifier)
	if err != nil {
		return err
	}

	return applyUpdates(cmd, schemaKey, objectType)
}

func applyUpdates(cmd *cobra.Command, schemaKey string, objectType *api.ObjectType) error {
	request := api.ObjectTypeUpdateRequest{}

	if cmd.Flags().Changed("name") {
		request.Name = objectTypeName // use the new name
	} else {
		request.Name = objectType.Name
	}

	if cmd.Flags().Changed("description") {
		request.Description = objectTypeDescription
	} else {
		request.Description = objectType.Description
	}

	if cmd.Flags().Changed("icon-id") {
		// update the id
		request.IconID = &objectTypeIconID
		request.IconName = ""
	} else if cmd.Flags().Changed("icon-name") {
		// icon-name takes precedence over icon-id which has a default
		request.IconID = nil
		request.IconName = objectTypeIconName
	} else {
		// don't change a thing; send the existing object_type_id
		request.IconID = &objectType.Icon.ID
		request.IconName = ""
	}

	updatedObjectType, err := api.DefaultClient().UpdateObjectType(strconv.Itoa(objectType.ID), &request)
	if err != nil {
		return err
	}

	if cmd.Flags().Changed("parent-type-id") {
		updatedObjectType, err = api.DefaultClient().UpdateObjectTypeParent(strconv.Itoa(objectType.ID), &objectTypeParentObjectTypeID)
	} else if cmd.Flags().Changed("parent-type-name") {
		if objectTypeParentObjectTypeName == "none" {
			updatedObjectType, err = api.DefaultClient().UpdateObjectTypeParent(strconv.Itoa(objectType.ID), nil)
		} else {
			parentObjectType, err := api.DefaultClient().GetObjectTypeByNameFromSchemaKey(objectSchemaKey, objectTypeParentObjectTypeName)
			if err != nil {
				return err
			}
			updatedObjectType, err = api.DefaultClient().UpdateObjectTypeParent(strconv.Itoa(objectType.ID), &parentObjectType.ID)
		}
	}
	if err != nil {
		return err
	}

	fmt.Printf("\nUpdated ObjectType %s (%d) in ObjectSchema %s\n\n", updatedObjectType.Name, updatedObjectType.ID, objectSchemaKey)

	insightFormat.WriteObjectType("Key", objectSchemaKey, updatedObjectType)

	return nil
}
