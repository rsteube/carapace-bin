package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var selectLayoutCmd = &cobra.Command{
	Use:   "select-layout",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(selectLayoutCmd).Standalone()

	selectLayoutCmd.Flags().BoolS("E", "E", false, "TODO description")
	selectLayoutCmd.Flags().BoolS("n", "n", false, "TODO description")
	selectLayoutCmd.Flags().BoolS("o", "o", false, "TODO description")
	selectLayoutCmd.Flags().BoolS("p", "p", false, "TODO description")
	selectLayoutCmd.Flags().StringS("t", "t", "", "target-pane")
	rootCmd.AddCommand(selectLayoutCmd)
}
