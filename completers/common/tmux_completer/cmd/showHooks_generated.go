package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var showHooksCmd = &cobra.Command{
	Use:   "show-hooks",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(showHooksCmd).Standalone()

	showHooksCmd.Flags().StringS("t", "t", "", "target-pane")
	rootCmd.AddCommand(showHooksCmd)
}
