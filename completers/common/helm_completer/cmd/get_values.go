package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/helm_completer/cmd/action"
	"github.com/spf13/cobra"
)

var get_valuesCmd = &cobra.Command{
	Use:   "values",
	Short: "download the values file for a named release",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	get_valuesCmd.Flags().BoolP("all", "a", false, "dump all (computed) values")
	get_valuesCmd.Flags().StringP("output", "o", "table", "prints the output in the specified format. Allowed values: table, json, yaml")
	get_valuesCmd.Flags().Int("revision", 0, "get the named release with revision")
	getCmd.AddCommand(get_valuesCmd)

	carapace.Gen(get_valuesCmd).FlagCompletion(carapace.ActionMap{
		"output": carapace.ActionValues("table", "json", "yaml"),
	})

	carapace.Gen(get_valuesCmd).PositionalCompletion(
		action.ActionReleases(),
	)
}
