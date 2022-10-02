package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var themesCmd = &cobra.Command{
	Use:   "themes",
	Short: "Prints list of available themes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(themesCmd).Standalone()

	themesCmd.Flags().BoolP("help", "h", false, "Prints help information")
	rootCmd.AddCommand(themesCmd)
}
