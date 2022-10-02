package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/helm_completer/cmd/action"
	"github.com/spf13/cobra"
)

var historyCmd = &cobra.Command{
	Use:   "history",
	Short: "fetch release history",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	historyCmd.Flags().Int("max", 256, "maximum number of revision to include in history")
	historyCmd.Flags().StringP("output", "o", "table", "prints the output in the specified format. Allowed values: table, json, yaml")
	rootCmd.AddCommand(historyCmd)

	carapace.Gen(historyCmd).FlagCompletion(carapace.ActionMap{
		"output": carapace.ActionValues("table", "json", "yaml"),
	})

	carapace.Gen(historyCmd).PositionalCompletion(
		action.ActionReleases(),
	)
}
