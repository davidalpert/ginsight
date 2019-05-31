package config

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// CmdAdd adds a configuration value
var CmdAdd = &cobra.Command{
	Use:   "add myKey myValue",
	Short: "Add an Insight CLI configuration setting",
	Args:  cobra.ExactArgs(2),
	Run:   addConfigurationValue,
}

func addConfigurationValue(cmd *cobra.Command, args []string) {
	key := args[0]
	val := args[1]
	viper.Set(key, val)
	viper.WriteConfig()
}
