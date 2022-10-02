package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/adb"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install a Flutter app on an attached device",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(installCmd).Standalone()

	installCmd.Flags().String("device-timeout", "", "Time in seconds to wait for devices to attach.")
	installCmd.Flags().String("device-user", "", "Identifier number for a user or work profile on Android only.")
	installCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	installCmd.Flags().Bool("no-uninstall-only", false, "Uninstall the app if already on the device.")
	installCmd.Flags().Bool("uninstall-only", false, "Uninstall the app if already on the device.")
	rootCmd.AddCommand(installCmd)

	carapace.Gen(installCmd).FlagCompletion(carapace.ActionMap{
		"device-user": adb.ActionUsers(),
	})
}
