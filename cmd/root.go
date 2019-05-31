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

package cmd

import (
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	config "github.com/davidalpert/ginsight/cmd/config"
	create "github.com/davidalpert/ginsight/cmd/create"
	update "github.com/davidalpert/ginsight/cmd/update"
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	rootCmd.SetArgs(os.Args[1:]) // without prog
	err := rootCmd.Execute()
	exitIfErr(err)
}

func exitIfErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var cfgFile string
var debugSet bool
var cfgInitialized = false

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ginsight",
	Short: "A CLI wrapper and API Client for the JIRA Insight API",
	Long: `
Interact with the JIRA Insight API.
`,
	SilenceUsage:  true,
	SilenceErrors: true,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },}
}

func init() {
	// We initialize our config via this callback (rather than inlining
	// config initialization) so that it has the benefit of reading in
	// bound CLI flag arguments like '--config customConfigFile'.
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.ginsight.yaml)")
	rootCmd.PersistentFlags().BoolVar(&debugSet, "debug", false, "enable debug logging")

	// Cobra also supports local Flags(), which will only run
	// when this action is called directly.

	// Register subcommands
	rootCmd.AddCommand(config.Cmd)
	rootCmd.AddCommand(create.Cmd)
	rootCmd.AddCommand(update.Cmd)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgInitialized {
		return
	}

	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		exitIfErr(err)

		// Search config in home directory with name ".insight" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".ginsight")
	}

	viper.AutomaticEnv() // read in environment variables that match

	fmt.Println("Looking for config file:", viper.ConfigFileUsed())
	// If a config file is found, read it in; if not, silently continue
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}

	cfgInitialized = true // only need to initialize viper once
}
