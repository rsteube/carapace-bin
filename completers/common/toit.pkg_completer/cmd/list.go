package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all available packages",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listCmd).Standalone()
	listCmd.Flags().StringP("output", "o", "list", "Defines the output format (valid: 'list', 'json')")
	listCmd.Flags().BoolP("verbose", "v", false, "Show more information")
	rootCmd.AddCommand(listCmd)

	carapace.Gen(listCmd).FlagCompletion(carapace.ActionMap{
		"output": carapace.ActionValues("list", "json"),
	})
}
