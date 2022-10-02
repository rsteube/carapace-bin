package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var bashCompletionCmd = &cobra.Command{
	Use:   "bash-completion",
	Short: "Output command line shell completion setup scripts",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bashCompletionCmd).Standalone()

	bashCompletionCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	bashCompletionCmd.Flags().Bool("no-overwrite", false, "Causes the given shell completion setup script not to be overwritten if it already exists.")
	bashCompletionCmd.Flags().Bool("overwrite", false, "Causes the given shell completion setup script to be overwritten if it already exists.")
	rootCmd.AddCommand(bashCompletionCmd)
}
