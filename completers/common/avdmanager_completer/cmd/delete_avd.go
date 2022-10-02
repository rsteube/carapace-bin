package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/avdmanager_completer/cmd/action"
	"github.com/spf13/cobra"
)

var delete_avdCmd = &cobra.Command{
	Use:   "avd",
	Short: "Deletes an Android Virtual Device",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(delete_avdCmd).Standalone()

	delete_avdCmd.Flags().StringP("name", "n", "", "Name of the AVD to delete.")
	deleteCmd.AddCommand(delete_avdCmd)

	carapace.Gen(delete_avdCmd).FlagCompletion(carapace.ActionMap{
		"name": action.ActionAvds(),
	})
}
