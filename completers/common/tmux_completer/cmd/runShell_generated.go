package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var runShellCmd = &cobra.Command{
	Use:   "run-shell",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(runShellCmd).Standalone()

	runShellCmd.Flags().BoolS("C", "C", false, "TODO description")
	runShellCmd.Flags().BoolS("b", "b", false, "TODO description")
	runShellCmd.Flags().StringS("d", "d", "", "delay")
	runShellCmd.Flags().StringS("t", "t", "", "target-pane")
	rootCmd.AddCommand(runShellCmd)
}
