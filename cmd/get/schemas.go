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
	"github.com/spf13/cobra"

	api "github.com/davidalpert/ginsight/api"
	insightFormat "github.com/davidalpert/ginsight/format"
)

// SchemasCmd represents the get command
var CmdGetSchemas = &cobra.Command{
	Use:   "schemas",
	Short: "Get information about your JIRA Insight object schemas",
	Long: `
Retreives a list of ObjectSchemas from the Insight API.
`,
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {

		objectSchemas, err := api.DefaultClient().GetObjectSchemas()
		if err != nil {
			return err
		}

		insightFormat.WriteObjectSchemasAsTable(objectSchemas)

		return nil
	},
}

func init() {
}
