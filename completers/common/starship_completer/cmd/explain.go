package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var explainCmd = &cobra.Command{
	Use:   "explain",
	Short: "Explains the currently showing modules",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(explainCmd).Standalone()

	explainCmd.Flags().BoolP("help", "h", false, "Prints help information")
	explainCmd.Flags().BoolP("version", "V", false, "Prints version information")
	rootCmd.AddCommand(explainCmd)
}
