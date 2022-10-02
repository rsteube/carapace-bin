package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var refreshClientCmd = &cobra.Command{
	Use:   "refresh-client",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(refreshClientCmd).Standalone()

	refreshClientCmd.Flags().StringS("A", "A", "", "pane:state")
	refreshClientCmd.Flags().StringS("B", "B", "", "name:what:format")
	refreshClientCmd.Flags().StringS("C", "C", "", "XxY")
	refreshClientCmd.Flags().BoolS("D", "D", false, "TODO description")
	refreshClientCmd.Flags().BoolS("L", "L", false, "TODO description")
	refreshClientCmd.Flags().BoolS("R", "R", false, "TODO description")
	refreshClientCmd.Flags().BoolS("S", "S", false, "TODO description")
	refreshClientCmd.Flags().BoolS("U", "U", false, "TODO description")
	refreshClientCmd.Flags().BoolS("c", "c", false, "TODO description")
	refreshClientCmd.Flags().StringS("f", "f", "", "flags")
	refreshClientCmd.Flags().BoolS("l", "l", false, "TODO description")
	refreshClientCmd.Flags().StringS("t", "t", "", "target-client")
	rootCmd.AddCommand(refreshClientCmd)
}
