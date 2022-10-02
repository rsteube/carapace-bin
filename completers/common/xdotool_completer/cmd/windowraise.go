package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var windowraiseCmd = &cobra.Command{
	Use:   "windowraise",
	Short: "Raise the window to the top of the stack",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(windowraiseCmd).Standalone()

	rootCmd.AddCommand(windowraiseCmd)
}
