package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/vagrant_completer/cmd/action"
	"github.com/spf13/cobra"
)

var suspendCmd = &cobra.Command{
	Use:   "suspend",
	Short: "suspends the machine",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(suspendCmd).Standalone()

	suspendCmd.Flags().BoolP("all-global", "a", false, "Suspend all running vms globally.")
	rootCmd.AddCommand(suspendCmd)

	carapace.Gen(suspendCmd).PositionalCompletion(
		action.ActionMachines(),
	)
}
