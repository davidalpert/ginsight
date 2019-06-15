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

package icon

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"

	api "github.com/davidalpert/ginsight/api"
	format "github.com/davidalpert/ginsight/format"
)

// SchemasCmd represents the get command
var CmdIconGet = &cobra.Command{
	Use:   "get identifier",
	Short: "Get information about one JIRA Insight ObjectIcon",
	Long: `
Retreives one ObjectIcon from the Insight API by name or id.

'identifier' can be either IconID or a Name
`,
	Args: cobra.ExactArgs(1),
	//PreRun: get
	RunE: iconGetE,
	Example: `
  # Get an icon by id
  ginsight icon get 42
  
  # Get a global icon by name
  ginsight icon get building --global
  
  # Get a schema icon by name
  ginsight icon get building --schema IT
  `,
}

func init() {}

func iconGetE(cmd *cobra.Command, args []string) error {
	iconIdentifier := args[0]

	if _, err := strconv.Atoi(iconIdentifier); err != nil {
		return getObjectIconByName(iconIdentifier)
	}

	return getObjectIconByID(iconIdentifier)
}

func getObjectIconByName(iconIdentifier string) (err error) {
	var foundIcon *api.ObjectIcon
	var scope string

	if global == true {
		fmt.Printf("Looking up global ObjectIcon by Name '%s'\n", iconIdentifier)
		foundIcon, err = api.DefaultClient().GetGlobalIconByName(iconIdentifier)
		scope = "global"
	} else {
		fmt.Printf("Looking up ObjectIcon by Name '%s' in ObjectSchema '%s'\n", iconIdentifier, objectSchemaKey)
		foundIcon, err = api.DefaultClient().GetSchemaIconByName(objectSchemaKey, iconIdentifier)
		scope = objectSchemaKey
	}
	if err != nil {
		return err
	}
	if foundIcon == nil {
		return fmt.Errorf("Could not find icon '%s' in '%s' scope", iconIdentifier, scope)
	}

	format.WriteObjectIcon("Key", scope, foundIcon)
	return nil
}

func getObjectIconByID(iconIdentifier string) error {
	fmt.Printf("Looking up ObjectIcon by ID '%s'\n", iconIdentifier)

	icon, err := api.DefaultClient().GetIconByID(iconIdentifier)
	if err != nil {
		return err
	}

	format.WriteObjectIcon("Key", "global", icon)
	return nil
}
