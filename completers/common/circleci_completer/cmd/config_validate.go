package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var config_validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Check that the config file is well formed.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_validateCmd).Standalone()
	config_validateCmd.PersistentFlags().StringP("config", "c", ".circleci/config.yml", "path to config file")
	config_validateCmd.Flags().StringP("org-slug", "o", "", "organization slug (for example: github/example-org), used when a config depends on private orbs belonging to that org")
	configCmd.AddCommand(config_validateCmd)

	carapace.Gen(config_validateCmd).FlagCompletion(carapace.ActionMap{
		"config": carapace.ActionFiles(),
	})

	carapace.Gen(config_validateCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
