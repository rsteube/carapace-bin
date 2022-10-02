package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var gui_versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display the currently running version of git gui",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gui_versionCmd).Standalone()

	guiCmd.AddCommand(gui_versionCmd)
}
