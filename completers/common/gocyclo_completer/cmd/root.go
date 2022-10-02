package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gocyclo",
	Short: "Calculate cyclomatic complexities of Go functions",
	Long:  "https://github.com/fzipp/gocyclo",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("avg", "avg", false, "show the average complexity over all functions")
	rootCmd.Flags().BoolS("avg-short", "avg-short", false, "show the average complexity over all functions without a label")
	rootCmd.Flags().StringS("ignore", "ignore", "", "exclude files matching the given regular expression")
	rootCmd.Flags().StringS("over", "over", "", "show functions with complexity > N only")
	rootCmd.Flags().StringS("top", "top", "", "show the top N most complex functions only")
	rootCmd.Flags().BoolS("total", "total", false, "show the total complexity for all functions")
	rootCmd.Flags().BoolS("total-short", "total-short", false, "show the total complexity for all functions without a label")

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
