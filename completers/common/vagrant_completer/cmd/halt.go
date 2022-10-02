package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/vagrant_completer/cmd/action"
	"github.com/spf13/cobra"
)

var haltCmd = &cobra.Command{
	Use:   "halt",
	Short: "stops the vagrant machine",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(haltCmd).Standalone()

	haltCmd.Flags().BoolP("force", "f", false, "Force shut down (equivalent of pulling power)")
	rootCmd.AddCommand(haltCmd)

	carapace.Gen(haltCmd).PositionalCompletion(
		action.ActionMachines(),
	)
}
