package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/helm_completer/cmd/action"
	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "display the status of the named release",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	statusCmd.Flags().StringP("output", "o", "table", "prints the output in the specified format. Allowed values: table, json, yaml")
	statusCmd.Flags().Int("revision", 0, "if set, display the status of the named release with revision")
	statusCmd.Flags().Bool("show-desc", false, "if set, display the description message of the named release")
	rootCmd.AddCommand(statusCmd)

	carapace.Gen(statusCmd).FlagCompletion(carapace.ActionMap{
		"output": carapace.ActionValues("table", "json", "yaml"),
	})

	carapace.Gen(statusCmd).PositionalCompletion(
		action.ActionReleases(),
	)
}
