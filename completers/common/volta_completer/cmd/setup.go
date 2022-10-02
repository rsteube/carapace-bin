package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Enables Volta for the current user / shell",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(setupCmd).Standalone()

	setupCmd.Flags().BoolP("help", "h", false, "Prints help information")
	setupCmd.Flags().Bool("quiet", false, "Prevents unnecessary output")
	setupCmd.Flags().Bool("verbose", false, "Enables verbose diagnostics")
	rootCmd.AddCommand(setupCmd)
}
