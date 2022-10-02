package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "swaybar",
	Short: "bar for swaywm",
	Long:  "https://github.com/swaywm/sway",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringP("bar_id", "b", "", "Bar ID for which to get the configuration.")
	rootCmd.Flags().BoolP("debug", "d", false, "Enable debugging.")
	rootCmd.Flags().BoolP("help", "h", false, "Show help message and quit.")
	rootCmd.Flags().StringP("socket", "s", "", "Connect to sway via socket.")
	rootCmd.Flags().BoolP("version", "v", false, "Show the version number and quit.")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"socket": carapace.ActionFiles(),
	})
}
