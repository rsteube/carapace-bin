package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var setWindowoptionCmd = &cobra.Command{
	Use:   "set-windowoption",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(setWindowoptionCmd).Standalone()

	setWindowoptionCmd.Flags().BoolS("F", "F", false, "TODO description")
	setWindowoptionCmd.Flags().BoolS("a", "a", false, "TODO description")
	setWindowoptionCmd.Flags().BoolS("g", "g", false, "TODO description")
	setWindowoptionCmd.Flags().BoolS("o", "o", false, "TODO description")
	setWindowoptionCmd.Flags().BoolS("q", "q", false, "TODO description")
	setWindowoptionCmd.Flags().StringS("t", "t", "", "target-window")
	setWindowoptionCmd.Flags().BoolS("u", "u", false, "TODO description")
	rootCmd.AddCommand(setWindowoptionCmd)
}
