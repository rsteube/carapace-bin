package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/vagrant_completer/cmd/action"
	"github.com/spf13/cobra"
)

var snapshot_popCmd = &cobra.Command{
	Use:   "pop",
	Short: "Restore state that was pushed onto the snapshot stack",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(snapshot_popCmd).Standalone()

	snapshot_popCmd.Flags().Bool("no-delete", false, "Don't delete the snapshot after the restore")
	snapshot_popCmd.Flags().Bool("no-provision", false, "Disable provisioning")
	snapshot_popCmd.Flags().Bool("no-start", false, "Don't start the snapshot after the restore")
	snapshot_popCmd.Flags().Bool("provision", false, "Enable provisioning")
	snapshot_popCmd.Flags().String("provision-with", "", "Enable only certain provisioners, by type or by name.")
	snapshotCmd.AddCommand(snapshot_popCmd)

	carapace.Gen(snapshot_popCmd).FlagCompletion(carapace.ActionMap{
		"provision-with": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return action.ActionProvisioners().Invoke(c).Filter(c.Parts).ToA()
		}),
	})

	carapace.Gen(snapshot_popCmd).PositionalCompletion(
		action.ActionLocalMachines(),
	)
}
