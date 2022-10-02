package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var getDesktopCmd = &cobra.Command{
	Use:   "get_desktop",
	Short: "Output the current desktop in view",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(getDesktopCmd).Standalone()

	rootCmd.AddCommand(getDesktopCmd)
}
