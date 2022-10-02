package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var setHookCmd = &cobra.Command{
	Use:   "set-hook",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(setHookCmd).Standalone()

	setHookCmd.Flags().StringS("t", "t", "", "target-pane")
	rootCmd.AddCommand(setHookCmd)
}
