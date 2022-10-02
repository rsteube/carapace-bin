package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Fetch the newest version of Homebrew and all formulae",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(updateCmd).Standalone()

	updateCmd.Flags().BoolP("debug", "d", false, "Display a trace of all shell commands as they are executed")
	updateCmd.Flags().BoolP("force", "f", false, "Always do a slower, full update check")
	updateCmd.Flags().BoolP("help", "h", false, "Show this message")
	updateCmd.Flags().Bool("merge", false, "Use git merge to apply updates")
	updateCmd.Flags().Bool("preinstall", false, "Run on auto-updates")
	updateCmd.Flags().BoolP("verbose", "v", false, "Print the directories checked and git operations performed")
	rootCmd.AddCommand(updateCmd)
}
