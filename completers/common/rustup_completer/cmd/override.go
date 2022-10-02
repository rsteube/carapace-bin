package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var overrideCmd = &cobra.Command{
	Use:   "override",
	Short: "Modify directory toolchain overrides",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(overrideCmd).Standalone()

	overrideCmd.Flags().BoolP("help", "h", false, "Prints help information")
	rootCmd.AddCommand(overrideCmd)
}
