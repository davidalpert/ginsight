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
	"regexp"

	"github.com/spf13/cobra"

	api "github.com/davidalpert/ginsight/api"
	"github.com/davidalpert/ginsight/format"
)

// SchemasCmd represents the get command
var CmdIconList = &cobra.Command{
	Use:   "list",
	Short: "List Icons from a JIRA Insight installation",
	Long: `
Retreives a list of Icons from the Insight API.
`,
	Args: cobra.NoArgs,
	RunE: iconsListE,
	Example: `
  # List all global icons 
  ginsight icon list --global
  
  # List all icons in a schema
  ginsight icon list --schema IT

  # Filter list with a text match 'ing'
  ginsight icon list --global --filter ing

  " Filter list with a case insensitive regex 'build'
  ginsight icon list --global --filter '(?i)build'
  `,
}

var iconNameFilter string

func init() {
	CmdIconList.Flags().StringVarP(&iconNameFilter, "filter", "f", "", "filter search results")
}

func iconsListE(cmd *cobra.Command, args []string) (err error) {
	var icons *[]api.ObjectIcon
	var schemaTag string

	if global {
		fmt.Printf("Looking up global Icons...\n")
		icons, err = api.DefaultClient().GetGlobalIcons()
		schemaTag = "global"
	} else {
		fmt.Printf("Looking up Icons for ObjectSchema '%s' ...\n", objectSchemaKey)
		icons, err = api.DefaultClient().GetSchemaIcons(objectSchemaKey)
		schemaTag = objectSchemaKey
	}

	if err != nil {
		return err
	}

	if iconNameFilter == "" {
		format.WriteObjectIcons("Key", schemaTag, icons)
	} else {
		filteredIcons := []api.ObjectIcon{}
		for _, icon := range *icons {
			if matched, _ := regexp.MatchString(iconNameFilter, icon.Name); matched {
				filteredIcons = append(filteredIcons, icon)
			}
		}
		format.WriteObjectIcons("Key", schemaTag, &filteredIcons)
	}

	return nil
}
