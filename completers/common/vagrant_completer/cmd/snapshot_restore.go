package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/vagrant_completer/cmd/action"
	"github.com/spf13/cobra"
)

var snapshot_restoreCmd = &cobra.Command{
	Use:   "restore",
	Short: "Restore a snapshot taken previously with snapshot save",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(snapshot_restoreCmd).Standalone()

	snapshot_restoreCmd.Flags().Bool("no-provision", false, "Disable provisioning")
	snapshot_restoreCmd.Flags().Bool("no-start", false, "Don't start the snapshot after the restore")
	snapshot_restoreCmd.Flags().Bool("provision", false, "Enable provisioning")
	snapshot_restoreCmd.Flags().String("provision-with", "", "Enable only certain provisioners, by type or by name.")
	snapshotCmd.AddCommand(snapshot_restoreCmd)

	carapace.Gen(snapshot_restoreCmd).FlagCompletion(carapace.ActionMap{
		"provision-with": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return action.ActionProvisioners().Invoke(c).Filter(c.Parts).ToA()
		}),
	})

	carapace.Gen(snapshot_restoreCmd).PositionalCompletion(
		action.ActionLocalMachines(),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionSnapshots(c.Args[0])
		}),
	)
}
