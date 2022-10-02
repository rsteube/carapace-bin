package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/minikube_completer/cmd/action"
	"github.com/spf13/cobra"
)

var config_defaultsCmd = &cobra.Command{
	Use:   "defaults",
	Short: "Lists all valid default values for PROPERTY_NAME",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	config_defaultsCmd.Flags().String("output", "table", "Output format. Accepted values: [json]")
	configCmd.AddCommand(config_defaultsCmd)

	carapace.Gen(config_defaultsCmd).FlagCompletion(carapace.ActionMap{
		"output": carapace.ActionValues("json"),
	})

	carapace.Gen(config_defaultsCmd).PositionalCompletion(
		action.ActionConfigNames(),
	)
}
