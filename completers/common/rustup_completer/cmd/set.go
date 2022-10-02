package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Alter rustup settings",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(setCmd).Standalone()

	setCmd.Flags().BoolP("help", "h", false, "Prints help information")
	rootCmd.AddCommand(setCmd)
}
