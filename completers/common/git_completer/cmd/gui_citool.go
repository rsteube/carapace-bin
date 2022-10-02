package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var gui_citoolCmd = &cobra.Command{
	Use:   "citool",
	Short: "Start git gui and arrange to make exactly one commit before exiting and returning to the shell",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gui_citoolCmd).Standalone()
	gui_citoolCmd.Flags().Bool("amend", false, "Automatically enter the Amend Last Commit mode of the interface")
	gui_citoolCmd.Flags().Bool("nocommit", false, "Behave as normal citool, but instead of making a commit simply terminate with a zero exit code")

	guiCmd.AddCommand(gui_citoolCmd)
}
