package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "adb",
	Short: "Android Debug Bridge",
	Long:  "https://developer.android.com/studio/command-line/adb",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("H", "H", false, "name of adb server host [default=localhost]")
	rootCmd.Flags().StringS("L", "L", "", "listen on given socket for adb server [default=tcp:localhost:5037]")
	rootCmd.Flags().BoolS("P", "P", false, "port of adb server [default=5037]")
	rootCmd.Flags().BoolS("a", "a", false, "listen on all network interfaces, not just localhost")
	rootCmd.Flags().BoolS("d", "d", false, "use USB device (error if multiple devices connected)")
	rootCmd.Flags().BoolS("e", "e", false, "use TCP/IP device (error if multiple TCP/IP devices available)")
	rootCmd.Flags().StringS("s", "s", "", "use device with given serial (overrides $ANDROID_SERIAL)")
	rootCmd.Flags().StringS("t", "t", "", "use device with given transport id")
}
