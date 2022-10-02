package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var displayMessageCmd = &cobra.Command{
	Use:   "display-message",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(displayMessageCmd).Standalone()

	displayMessageCmd.Flags().StringS("F", "F", "", "format")
	displayMessageCmd.Flags().BoolS("I", "I", false, "TODO description")
	displayMessageCmd.Flags().BoolS("N", "N", false, "TODO description")
	displayMessageCmd.Flags().BoolS("a", "a", false, "TODO description")
	displayMessageCmd.Flags().StringS("c", "c", "", "target-client")
	displayMessageCmd.Flags().StringS("d", "d", "", "delay")
	displayMessageCmd.Flags().BoolS("p", "p", false, "TODO description")
	displayMessageCmd.Flags().StringS("t", "t", "", "target-pane")
	displayMessageCmd.Flags().BoolS("v", "v", false, "TODO description")
	rootCmd.AddCommand(displayMessageCmd)
}
