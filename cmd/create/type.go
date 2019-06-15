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

var CmdCreateType = &cobra.Command{
	Use:   "type",
	Short: "Create an ObjectType in a JIRA Insight ObjectSchema",
	Long: `
Creates a new ObjectType in a JIRA Insight ObjectSchema
`,

	Args: cobra.ExactArgs(1),
	Example: `
  # create a new ObjectType in the IT schema with the default icon
  insight create type --key IT --name Vendor

  # create a new ObjectType in the IT schema with a description
  insight create type --key IT --name Vendor --description "A vendor provides products or services"

  # create a new ObjectType in the IT schema with an icon by ID
  insight create type --key IT --name Vendor --icon-id 42

  # create a new ObjectType in the IT schema with an icon by name
  insight create type --key IT --name Vendor --icon-name cloud
`,
	RunE: createType,
}

func init() {
	CmdCreateType.Flags().StringVar(&objectTypeDescription, "description", "", "description of the object type to create")
	CmdCreateType.Flags().StringVarP(&objectTypeParentObjectTypeName, "parent-type-name", "p", "", "name of the parent object type")
	CmdCreateType.Flags().IntVarP(&objectTypeParentObjectTypeID, "parent-type-id", "P", -1, "id of the parent object type")
	CmdCreateType.Flags().IntVarP(&objectTypeIconID, "icon-id", "i", 1, "id of the icon")     // icon is required, 1-indexed, so 1 is the default iconId
	CmdCreateType.Flags().StringVar(&objectTypeIconName, "icon-name", "", "name of the icon") // icon will be looked up against the global set
}

func createType(cmd *cobra.Command, args []string) error {
	if cmd.Flags().Changed("parent-type-name") && cmd.Flags().Changed("parent-type-id") {
		return fmt.Errorf("Cannot specify the parent type by both name and id; please pick one method to specify the parent type.")
	}

	objectTypeName = args[0]

	fmt.Printf("Looking up ObjectSchema by key: %s ... ", objectSchemaKey)
	client := api.DefaultClient()

	schema, err := client.GetObjectSchemaByKey(objectSchemaKey)
	if err != nil {
		return err
	}

	fmt.Printf("found it!\n")

	if cmd.Flags().Changed("parent-type-name") {
		fmt.Printf("Looking up ObjectType by name '%s' in schema %s (%d) ... ", objectTypeParentObjectTypeName, objectSchemaKey, schema.ID)

		parentObjectTypes, err := client.GetObjectTypesByNameFromSchemaID(strconv.Itoa(schema.ID), objectTypeParentObjectTypeName)
		if err != nil {
			return err
		}

		if len(*parentObjectTypes) != 1 {
			return &api.MultipleObjectTypesFoundError{
				SchemaID:       schema.Key,
				ObjectTypeName: objectTypeParentObjectTypeName,
				FoundTypes:     parentObjectTypes,
			}
		}

		objectTypeParentObjectTypeID = (*parentObjectTypes)[0].ID
		fmt.Printf("found the id: %d\n", objectTypeParentObjectTypeID)
	}

	typeCreate := api.ObjectTypeCreateRequest{
		ObjectSchemaID: schema.ID,
		Name:           objectTypeName,
		Description:    objectTypeDescription,
		IconID:         &objectTypeIconID,
		IconName:       objectTypeIconName,
	}

	if cmd.Flags().Changed("icon-name") {
		// icon-name takes precedence over icon-id which has a default
		fmt.Printf("icon-name has changed, using %s over %d", typeCreate.IconName, typeCreate.IconID)
		typeCreate.IconID = nil
	}

	if objectTypeParentObjectTypeID >= 0 {
		typeCreate.ParentObjectTypeID = objectTypeParentObjectTypeID
	}

	fmt.Printf("Creating new ObjectType: %v\n", typeCreate)

	objectType, err := client.CreateObjectType(&typeCreate)
	if err != nil {
		return err
	}

	fmt.Printf("\nCreated ObjectType %s (%d) in ObjectSchema %s\n\n", objectType.Name, objectType.ID, objectSchemaKey)

	insightFormat.WriteObjectType("Key", objectSchemaKey, objectType)
	return nil
}
