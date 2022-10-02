package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var forceUnlockCmd = &cobra.Command{
	Use:   "force-unlock",
	Short: "Release a stuck lock on the current workspace",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(forceUnlockCmd).Standalone()

	forceUnlockCmd.Flags().BoolS("force", "force", false, "Don't ask for input for unlock confirmation.")
	rootCmd.AddCommand(forceUnlockCmd)

	// TODO lock_id positional completion
}
