package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var search_repoCmd = &cobra.Command{
	Use:   "repo",
	Short: "search repositories for a keyword in charts",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	search_repoCmd.Flags().Bool("devel", false, "use development versions (alpha, beta, and release candidate releases), too. Equivalent to version '>0.0.0-0'. If --version is set, this is ignored")
	search_repoCmd.Flags().Uint("max-col-width", 50, "maximum column width for output table")
	search_repoCmd.Flags().StringP("output", "o", "table", "prints the output in the specified format. Allowed values: table, json, yaml")
	search_repoCmd.Flags().BoolP("regexp", "r", false, "use regular expressions for searching repositories you have added")
	search_repoCmd.Flags().String("version", "", "search using semantic versioning constraints on repositories you have added")
	search_repoCmd.Flags().BoolP("versions", "l", false, "show the long listing, with each version of each chart on its own line, for repositories you have added")
	searchCmd.AddCommand(search_repoCmd)

	carapace.Gen(search_repoCmd).FlagCompletion(carapace.ActionMap{
		"output": carapace.ActionValues("table", "json", "yaml"),
	})
}
