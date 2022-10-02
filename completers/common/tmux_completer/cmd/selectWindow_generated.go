package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var selectWindowCmd = &cobra.Command{
	Use:   "select-window",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(selectWindowCmd).Standalone()

	selectWindowCmd.Flags().BoolS("T", "T", false, "TODO description")
	selectWindowCmd.Flags().BoolS("l", "l", false, "TODO description")
	selectWindowCmd.Flags().BoolS("n", "n", false, "TODO description")
	selectWindowCmd.Flags().BoolS("p", "p", false, "TODO description")
	selectWindowCmd.Flags().StringS("t", "t", "", "target-window")
	rootCmd.AddCommand(selectWindowCmd)
}
