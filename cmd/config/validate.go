package config

import (
	"fmt"

	api "github.com/davidalpert/ginsight/api"
	"github.com/spf13/cobra"
)

// CmdValidate validates the Jira configuration by requesting info about the configured Jira credentials
var CmdValidate = &cobra.Command{
	Use:           "-- validate",
	Short:         "Validate the Insight Client configuration",
	Args:          cobra.NoArgs,
	RunE:          validateConfiguration,
	SilenceUsage:  true,
	SilenceErrors: true,
}

func validateConfiguration(cmd *cobra.Command, args []string) error {
	fmt.Println("Validating your Insight client configuration...\n")

	client := api.DefaultClient()
	me, err := client.GetMe()

	if err != nil {
		return err
	}

	fmt.Printf("Successfully authenticated to %s as '%s' (%s)\n", client.BaseURL, me.Name, me.DisplayName)
	return nil
}
