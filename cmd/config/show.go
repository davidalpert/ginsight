package config

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// CmdShow shows a configuration value
var CmdShow = &cobra.Command{
	Use:   "show",
	Short: "Show an Insight CLI configuration setting",
	Args:  cobra.ExactArgs(1),
	Run:   writeConfiguration,
}

func writeConfiguration(cmd *cobra.Command, args []string) {
	key := args[0]
	fmt.Printf("- %s: %v\n", key, viper.Get(key))
}
