package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var clockModeCmd = &cobra.Command{
	Use:   "clock-mode",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(clockModeCmd).Standalone()

	clockModeCmd.Flags().StringS("t", "t", "", "target-pane")
	rootCmd.AddCommand(clockModeCmd)
}
