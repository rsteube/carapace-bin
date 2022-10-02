package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var config_processCmd = &cobra.Command{
	Use:   "process",
	Short: "Validate config and display expanded configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_processCmd).Standalone()
	config_processCmd.Flags().StringP("org-slug", "o", "", "organization slug (for example: github/example-org), used when a config depends on private orbs belonging to that org")
	config_processCmd.Flags().String("pipeline-parameters", "", "YAML/JSON map of pipeline parameters, accepts either YAML/JSON directly or file path (for example: my-params.yml)")
	configCmd.AddCommand(config_processCmd)

	carapace.Gen(config_processCmd).FlagCompletion(carapace.ActionMap{
		"pipeline-parameters": carapace.ActionFiles(),
	})

	carapace.Gen(config_processCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
