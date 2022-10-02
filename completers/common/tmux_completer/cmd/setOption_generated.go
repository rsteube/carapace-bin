package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var setOptionCmd = &cobra.Command{
	Use:   "set-option",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(setOptionCmd).Standalone()

	setOptionCmd.Flags().BoolS("F", "F", false, "TODO description")
	setOptionCmd.Flags().BoolS("U", "U", false, "TODO description")
	setOptionCmd.Flags().BoolS("a", "a", false, "TODO description")
	setOptionCmd.Flags().BoolS("g", "g", false, "TODO description")
	setOptionCmd.Flags().BoolS("o", "o", false, "TODO description")
	setOptionCmd.Flags().BoolS("p", "p", false, "TODO description")
	setOptionCmd.Flags().BoolS("q", "q", false, "TODO description")
	setOptionCmd.Flags().BoolS("s", "s", false, "TODO description")
	setOptionCmd.Flags().StringS("t", "t", "", "target-pane")
	setOptionCmd.Flags().BoolS("u", "u", false, "TODO description")
	setOptionCmd.Flags().BoolS("w", "w", false, "TODO description")
	rootCmd.AddCommand(setOptionCmd)
}
