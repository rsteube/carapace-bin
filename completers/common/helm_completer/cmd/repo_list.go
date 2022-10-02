package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var repo_listCmd = &cobra.Command{
	Use:   "list",
	Short: "list chart repositories",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	repo_listCmd.Flags().StringP("output", "o", "table", "prints the output in the specified format. Allowed values: table, json, yaml")
	repoCmd.AddCommand(repo_listCmd)

	carapace.Gen(repo_listCmd).FlagCompletion(carapace.ActionMap{
		"output": carapace.ActionValues("table", "json", "yaml"),
	})
}
