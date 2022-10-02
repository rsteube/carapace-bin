package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var getwindowfocusCmd = &cobra.Command{
	Use:   "getwindowfocus",
	Short: "Prints the window id of the currently focused window",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(getwindowfocusCmd).Standalone()
	getwindowfocusCmd.Flags().BoolS("f", "f", false, "ignore WM_CLASS property")

	rootCmd.AddCommand(getwindowfocusCmd)
}
