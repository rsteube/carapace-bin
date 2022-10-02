package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var copyModeCmd = &cobra.Command{
	Use:   "copy-mode",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(copyModeCmd).Standalone()

	copyModeCmd.Flags().StringS("s", "s", "", "src-pane")
	copyModeCmd.Flags().StringS("t", "t", "", "target-pane")
	rootCmd.AddCommand(copyModeCmd)
}
