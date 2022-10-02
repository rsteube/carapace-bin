package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var swapWindowCmd = &cobra.Command{
	Use:   "swap-window",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(swapWindowCmd).Standalone()

	swapWindowCmd.Flags().BoolS("d", "d", false, "TODO description")
	swapWindowCmd.Flags().StringS("s", "s", "", "src-window")
	swapWindowCmd.Flags().StringS("t", "t", "", "dst-window")
	rootCmd.AddCommand(swapWindowCmd)
}
