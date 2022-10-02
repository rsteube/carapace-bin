package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/adb"
	"github.com/spf13/cobra"
)

var attachCmd = &cobra.Command{
	Use:   "attach",
	Short: "Attach to a running app",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(attachCmd).Standalone()

	attachCmd.Flags().String("app-id", "", "The package name (Android) or bundle identifier (iOS) for the app")
	attachCmd.Flags().StringArray("dart-define", []string{}, "Additional key-value pairs that will be available as constants.")
	attachCmd.Flags().String("dds-port", "", "When this value is provided, the Dart Development Service (DDS) will be bound to the provided port.")
	attachCmd.Flags().Bool("debug", false, "Build a debug version of your app (default mode).")
	attachCmd.Flags().String("debug-uri", "", "The URL at which the observatory is listening.")
	attachCmd.Flags().String("device-timeout", "", "Time in seconds to wait for devices to attach.")
	attachCmd.Flags().String("device-user", "", "Identifier number for a user or work profile on Android only.")
	attachCmd.Flags().String("device-vmservice-port", "", "Look for vmservice connections only from the specified port.")
	attachCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	attachCmd.Flags().String("host-vmservice-port", "", "When a device-side vmservice port is forwarded to a host-side port, use this value as the host port.")
	attachCmd.Flags().Bool("no-null-assertions", false, "Do not perform additional null assertions on the boundaries of migrated and un-migrated code.")
	attachCmd.Flags().Bool("no-track-widget-creation", false, "Do not track widget creation locations.")
	attachCmd.Flags().Bool("null-assertions", false, "Perform additional null assertions on the boundaries of migrated and un-migrated code.")
	attachCmd.Flags().String("pid-file", "", "Specify a file to write the process ID to.")
	attachCmd.Flags().Bool("profile", false, "Build a version of your app specialized for performance profiling.")
	attachCmd.Flags().StringP("target", "t", "", "The main entry-point file of the application, as run on the device.")
	attachCmd.Flags().Bool("track-widget-creation", false, "Track widget creation locations.")
	rootCmd.AddCommand(attachCmd)

	carapace.Gen(attachCmd).FlagCompletion(carapace.ActionMap{
		"app-id":      adb.ActionPackages(), // TODO ios
		"device-user": adb.ActionUsers(),
		"pid-file":    carapace.ActionFiles(),
		"target":      carapace.ActionFiles(".dart"),
	})
}
