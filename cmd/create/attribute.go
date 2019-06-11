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
	"strings"

	"github.com/spf13/cobra"

	api "github.com/davidalpert/ginsight/api"
	insightFormat "github.com/davidalpert/ginsight/format"
)

var attributeDescription string
var attributeType string
var attributeDefaultType string

var CmdCreateTypeAttribute = &cobra.Command{
	Use:   "attribute",
	Short: "Create an ObjectTypeAttribute in a JIRA Insight ObjectSchema",
	Long: `
Creates a new ObjectTypeAttribute in a JIRA Insight ObjectSchema
`,

	Args: cobra.ExactArgs(2), // typename attributename
	Example: `
  # create a new ObjectTypeAttribute in the IT schema
  insight create attribute Datacentre Address1 --attr-type default --default-type text --description "Address Line 1" --key IT
`,
	PreRunE: createAttributePreRunE,
	RunE:    createAttribute,
}

func init() {
	CmdCreateTypeAttribute.Flags().StringVar(&attributeType, "attr-type", "", "name of the attribute type")
	CmdCreateTypeAttribute.MarkFlagRequired("attr-type")

	CmdCreateTypeAttribute.Flags().StringVar(&attributeDefaultType, "default-type", "", "name of the default type (")
	CmdCreateTypeAttribute.Flags().StringVar(&objectTypeDescription, "description", "", "description of the object type to create")
}

func createAttributePreRunE(cmd *cobra.Command, args []string) error {
	if strings.ToLower(attributeType) == "default" && cmd.Flags().Changed("default-type") == false {
		return fmt.Errorf("When creating an Attribute of type 'default' you must specify '--default-type")
	}
	return nil
}

func createAttribute(cmd *cobra.Command, args []string) error {
	objectTypeIdentifier := args[0]
	attributeName := args[1]

	fmt.Printf("Creating attribute '%s' on type '%s' in schema '%s'\n", attributeName, objectTypeIdentifier, objectSchemaKey)

	var objectType *api.ObjectType
	if _, atoiErr := strconv.Atoi(objectTypeIdentifier); atoiErr == nil {
		fmt.Printf("Looking up ObjectType ID '%s' in schema '%s'\n", objectTypeIdentifier, objectSchemaKey)
		foundType, objectTypeErr := api.DefaultClient().GetObjectTypeByID(objectTypeIdentifier)
		if objectTypeErr != nil {
			return objectTypeErr
		}
		objectType = foundType
	} else {
		fmt.Printf("Looking up ObjectType ID for '%s' in schema '%s'\n", objectTypeIdentifier, objectSchemaKey)
		foundTypes, objectTypesErr := api.DefaultClient().GetObjectTypesByNameFromSchemaKey(objectSchemaKey, objectTypeIdentifier)
		if objectTypesErr != nil {
			return objectTypesErr
		}
		if len(*foundTypes) > 1 {
			return &api.MultipleObjectTypesFoundError{
				SchemaID:       objectSchemaKey,
				ObjectTypeName: objectTypeIdentifier,
				FoundTypes:     foundTypes,
			}
		}
		objectType = &(*foundTypes)[0]
	}

	fmt.Printf("Creating the attribute on type %s (%d)\n", objectType.Name, objectType.ID)
	attributeTypeID := api.AttributeTypeNameToID(attributeType)
	switch attributeTypeID {
	case api.AttributeTypeIDsByName["default"]:
		return createDefaultAttribute(attributeName, objectType)
	case api.AttributeTypeIDsByName["user"]:
		return fmt.Errorf("don't know how to create attributes of type %s\n", attributeType)
	default:
		return fmt.Errorf("don't know how to create attributes of type %s\n", attributeType)
	}

	return fmt.Errorf("Creating attributes of type '%s' are not supported.", attributeType)
}

func createDefaultAttribute(attributeName string, objectType *api.ObjectType) error {
	createDefaultAttributeRequest := api.ObjectTypeDefaultAttributeCreateRequest{
		Name:          attributeName,
		Description:   attributeDescription,
		TypeID:        api.AttributeTypeIDsByName["default"],
		DefaultTypeID: api.AttributeDefaultTypeIDsByName[attributeDefaultType],
	}

	//fmt.Println(createDefaultAttributeRequest)

	attr, err := api.DefaultClient().CreateObjectTypeDefaultAttribute(strconv.Itoa(objectType.ID), &createDefaultAttributeRequest)
	if err != nil {
		return err
	}

	fmt.Printf("\nCreated Attribute %s on ObjectType %s (%d) in ObjectSchema %s\n\n", attr.Name, objectType.Name, objectType.ID, objectSchemaKey)

	insightFormat.WriteObjectTypeAttribute("Key", objectSchemaKey, objectType, attr)

	return nil
}
