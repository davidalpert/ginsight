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

	"github.com/spf13/cobra"

	api "github.com/davidalpert/ginsight/api"
	"github.com/davidalpert/ginsight/format"
)

// SchemasCmd represents the get command
var CmdIconGet = &cobra.Command{
	Use:   "list",
	Short: "List Icons from a JIRA Insight installation",
	Long: `
Retreives a list of Icons from the Insight API.
`,
	Args: cobra.NoArgs,
	RunE: iconsListE,
}

func iconsListE(cmd *cobra.Command, args []string) error {
	if global {
		fmt.Printf("Looking up global Icons...\n")
		icons, err := api.DefaultClient().GetGlobalIcons()
		if err != nil {
			return err
		}
		format.WriteIcons("Key", "global", icons)
	} else {
		fmt.Printf("Looking up Icons for ObjectSchema '%s' ...\n", objectSchemaKey)
		icons, err := api.DefaultClient().GetSchemaIcons(objectSchemaKey)
		if err != nil {
			return err
		}
		format.WriteIcons("Key", objectSchemaKey, icons)
	}

	return nil
}
