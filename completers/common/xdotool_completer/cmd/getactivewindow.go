package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var getactivewindowCmd = &cobra.Command{
	Use:   "getactivewindow",
	Short: "Output the current active window",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(getactivewindowCmd).Standalone()

	rootCmd.AddCommand(getactivewindowCmd)
}
