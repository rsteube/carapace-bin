package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var newSessionCmd = &cobra.Command{
	Use:   "new-session",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(newSessionCmd).Standalone()

	newSessionCmd.Flags().BoolS("A", "A", false, "TODO description")
	newSessionCmd.Flags().BoolS("D", "D", false, "TODO description")
	newSessionCmd.Flags().BoolS("E", "E", false, "TODO description")
	newSessionCmd.Flags().StringS("F", "F", "", "format")
	newSessionCmd.Flags().BoolS("P", "P", false, "TODO description")
	newSessionCmd.Flags().BoolS("X", "X", false, "TODO description")
	newSessionCmd.Flags().StringS("c", "c", "", "start-directory")
	newSessionCmd.Flags().BoolS("d", "d", false, "TODO description")
	newSessionCmd.Flags().StringS("e", "e", "", "environment")
	newSessionCmd.Flags().StringS("f", "f", "", "flags")
	newSessionCmd.Flags().StringS("n", "n", "", "window-name")
	newSessionCmd.Flags().StringS("s", "s", "", "session-name")
	newSessionCmd.Flags().StringS("t", "t", "", "target-session")
	newSessionCmd.Flags().StringS("x", "x", "", "width")
	newSessionCmd.Flags().StringS("y", "y", "", "height")
	rootCmd.AddCommand(newSessionCmd)
}
