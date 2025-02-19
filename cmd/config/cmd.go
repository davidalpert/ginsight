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

// Config subcommands for the ginsight CLI
package config

import (
	"github.com/spf13/cobra"
)

// config.Cmd represents the root config command
var Cmd = &cobra.Command{
	Use:   "config",
	Short: "Configure the Insight CLI",
	Long:  `Manage Insight CLI config settings`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if cmd.Flags().Changed("add") {
			CmdAdd.SetArgs(args)
			return CmdAdd.Execute()
		}

		if cmd.Flags().Changed("validate") {
			CmdValidate.SetArgs(args)
			return CmdValidate.Execute()
		}

		numArgs := len(args)
		if numArgs == 1 {
			CmdShow.SetArgs(args)
			return CmdShow.Execute()
		}

		cmd.Help()
		return nil
	},
}

func init() {
	Cmd.Flags().BoolP("add", "a", false, "name")
	Cmd.Flags().BoolP("validate", "", false, "validate your configuration")
}
