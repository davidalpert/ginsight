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

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var Cmd = &cobra.Command{
	Use:   "create",
	Short: "Create schema or objects in Insight",
	Long: `
Provides subcommands to create various schema elements or objects in a Jira Insight installation.

`,
	Example: `
# List all object schemas
insight create type --schema IT`,
	PersistentPreRunE: getPersistentPreRunE,
}

// used for flags
var objectSchemaKey = ""

func getPersistentPreRunE(cmd *cobra.Command, args []string) error {
	// populate and validate before invoking the subcommand
	objectSchemaKey = viper.GetString("schema")
	if objectSchemaKey == "" {
		return fmt.Errorf("must specify an ObjectSchema; use the --schema flag with a schema key or set the insight.schema config value.\n")
	}
	return nil
}

func init() {
	Cmd.PersistentFlags().String("schema", "", "key of a schema you'd like to see in detail")
	viper.BindPFlag("schema", Cmd.PersistentFlags().Lookup("schema"))

	// Subcommands
	Cmd.AddCommand(CmdCreateType)
	//Cmd.AddCommand(CmdCreateTypeAttribute)
}
