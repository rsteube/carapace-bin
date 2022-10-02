package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var ifShellCmd = &cobra.Command{
	Use:   "if-shell",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ifShellCmd).Standalone()

	ifShellCmd.Flags().BoolS("F", "F", false, "TODO description")
	ifShellCmd.Flags().BoolS("b", "b", false, "TODO description")
	ifShellCmd.Flags().StringS("t", "t", "", "target-pane")
	rootCmd.AddCommand(ifShellCmd)
}
