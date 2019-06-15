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
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var Cmd = &cobra.Command{
	Use:   "icon",
	Short: "Interact with Insight Icons",
	Long: `
Provides subcommands to interact with Icons in a Jira Insight installation.

`,
	Example: `
# List all icons 
insight icon get
`,

	PersistentPreRunE: iconPersistentPreRunE,
}

// used for flags
var objectSchemaKey = ""
var global = false

func iconPersistentPreRunE(cmd *cobra.Command, args []string) error {
	objectSchemaKey = viper.GetString("schema")
	global = viper.GetBool("global")

	if global {
		// have to clear any default schema key
		objectSchemaKey = ""
	} else if objectSchemaKey == "" {
		// default to global icon set
		global = true
	}

	return nil
}

func init() {
	Cmd.PersistentFlags().String("schema", "", "limit scope to a specific ObjectSchema key")
	viper.BindPFlag("schema", Cmd.PersistentFlags().Lookup("schema"))

	Cmd.PersistentFlags().Bool("global", false, "use global icon scope")
	viper.BindPFlag("global", Cmd.PersistentFlags().Lookup("global"))

	// Subcommands
	Cmd.AddCommand(CmdIconGet)
}
