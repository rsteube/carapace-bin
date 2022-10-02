package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/avdmanager_completer/cmd/action"
	"github.com/spf13/cobra"
)

var create_avdCmd = &cobra.Command{
	Use:   "avd",
	Short: "Creates a new Android Virtual Device",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(create_avdCmd).Standalone()

	create_avdCmd.Flags().StringP("abi", "b", "", "The ABI to use for the AVD.")
	create_avdCmd.Flags().StringP("device", "d", "", "The optional device definition to use.")
	create_avdCmd.Flags().BoolP("force", "f", false, "Forces creation (overwrites an existing AVD)")
	create_avdCmd.Flags().StringP("name", "n", "", "Name of the new AVD.")
	create_avdCmd.Flags().StringP("package", "k", "", "Package path of the system image for this AVD")
	create_avdCmd.Flags().StringP("path", "p", "", "Directory where the new AVD will be created.")
	create_avdCmd.Flags().StringP("sdcard", "c", "", "Path to a shared SD card image, or size of a new sdcard for the new AVD.")
	create_avdCmd.Flags().StringP("snapshot", "a", "", "Place a snapshots file in the AVD, to enable persistence.")
	create_avdCmd.Flags().StringP("tag", "g", "", "The sys-img tag to use for the AVD.")
	createCmd.AddCommand(create_avdCmd)

	carapace.Gen(create_avdCmd).FlagCompletion(carapace.ActionMap{
		"device":   action.ActionDevices(),
		"path":     carapace.ActionDirectories(),
		"sdcard":   carapace.ActionFiles(),
		"snapshot": carapace.ActionFiles(),
	})
}
