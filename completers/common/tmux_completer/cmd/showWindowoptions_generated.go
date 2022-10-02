package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var showWindowoptionsCmd = &cobra.Command{
	Use:   "show-windowoptions",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(showWindowoptionsCmd).Standalone()

	showWindowoptionsCmd.Flags().BoolS("g", "g", false, "TODO description")
	showWindowoptionsCmd.Flags().StringS("t", "t", "", "target-window")
	showWindowoptionsCmd.Flags().BoolS("v", "v", false, "TODO description")
	rootCmd.AddCommand(showWindowoptionsCmd)
}
