package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/avdmanager_completer/cmd/action"
	"github.com/spf13/cobra"
)

var move_avdCmd = &cobra.Command{
	Use:   "avd",
	Short: "Moves or renames an Android Virtual Device",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(move_avdCmd).Standalone()

	move_avdCmd.Flags().StringP("name", "n", "", "Name of the AVD to move or rename.")
	move_avdCmd.Flags().StringP("path", "p", "", "Path to the AVD's new directory.")
	move_avdCmd.Flags().StringP("rename", "r", "", "New name of the AVD.")
	moveCmd.AddCommand(move_avdCmd)

	carapace.Gen(move_avdCmd).FlagCompletion(carapace.ActionMap{
		"name": action.ActionAvds(),
		"path": carapace.ActionDirectories(),
	})
}
