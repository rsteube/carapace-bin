package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var customizeModeCmd = &cobra.Command{
	Use:   "customize-mode",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(customizeModeCmd).Standalone()

	customizeModeCmd.Flags().StringS("F", "F", "", "format")
	customizeModeCmd.Flags().StringS("f", "f", "", "filter")
	customizeModeCmd.Flags().StringS("t", "t", "", "target-pane")
	rootCmd.AddCommand(customizeModeCmd)
}
