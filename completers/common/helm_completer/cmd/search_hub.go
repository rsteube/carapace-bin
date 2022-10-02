package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var search_hubCmd = &cobra.Command{
	Use:   "hub",
	Short: "search for charts in the Artifact Hub or your own hub instance",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	search_hubCmd.Flags().String("endpoint", "https://hub.helm.sh", "Hub instance to query for charts")
	search_hubCmd.Flags().Uint("max-col-width", 50, "maximum column width for output table")
	search_hubCmd.Flags().StringP("output", "o", "table", "prints the output in the specified format. Allowed values: table, json, yaml")
	searchCmd.AddCommand(search_hubCmd)

	carapace.Gen(search_hubCmd).FlagCompletion(carapace.ActionMap{
		"output": carapace.ActionValues("table", "json", "yaml"),
	})
}
