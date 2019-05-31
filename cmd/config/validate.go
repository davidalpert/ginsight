package config

import (
	"fmt"

	insight "github.com/davidalpert/ginsight/insight"
	"github.com/spf13/cobra"
)

// CmdValidate validates the Jira configuration by requesting info about the configured Jira credentials
var CmdValidate = &cobra.Command{
	Use:   "-- validate",
	Short: "Validate the Insight Client configuration",
	Args:  cobra.NoArgs,
	RunE:  validateConfiguration,
}

func validateConfiguration(cmd *cobra.Command, args []string) error {
	fmt.Println("Validating your Insight client configuration...\n")

	client := insight.DefaultClient()
	me, err := client.GetMe()
	if err != nil {
		return err
	}

	fmt.Printf("Successfully authenticated to %s as '%s' (%s)\n", client.BaseURL, me.Name, me.DisplayName)
	return nil
}
