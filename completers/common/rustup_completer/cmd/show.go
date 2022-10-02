package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show the active and installed toolchains or profiles",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(showCmd).Standalone()

	showCmd.Flags().BoolP("help", "h", false, "Prints help information")
	rootCmd.AddCommand(showCmd)
}
