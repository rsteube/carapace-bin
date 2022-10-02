package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var detachClientCmd = &cobra.Command{
	Use:   "detach-client",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(detachClientCmd).Standalone()

	detachClientCmd.Flags().StringS("E", "E", "", "shell-command")
	detachClientCmd.Flags().BoolS("P", "P", false, "TODO description")
	detachClientCmd.Flags().BoolS("a", "a", false, "TODO description")
	detachClientCmd.Flags().StringS("s", "s", "", "target-session")
	detachClientCmd.Flags().StringS("t", "t", "", "target-client")
	rootCmd.AddCommand(detachClientCmd)
}
