package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/flutter_completer/cmd/action"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "flutter",
	Short: "Manage your Flutter app development",
	Long:  "https://flutter.dev/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringP("device-id", "d", "", "Target device id or name (prefixes allowed).")
	rootCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	rootCmd.Flags().Bool("suppress-analytics", false, "Suppress analytics reporting when this command runs.")
	rootCmd.Flags().BoolP("verbose", "v", false, "Noisy logging, including all shell commands executed.")
	rootCmd.Flags().Bool("version", false, "Reports the version of this tool.")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"device-id": action.ActionDevices(),
	})
}
