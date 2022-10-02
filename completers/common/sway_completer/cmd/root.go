package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "sway",
	Short: "An i3-compatible Wayland compositor",
	Long:  "https://swaywm.org/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringP("config", "c", "", "Specify a config file.")
	rootCmd.Flags().BoolP("debug", "d", false, "Enables full logging, including debug information.")
	rootCmd.Flags().Bool("get-socketpath", false, "Gets the IPC socket path and prints it, then exits.")
	rootCmd.Flags().BoolP("help", "h", false, "Show help message and quit.")
	rootCmd.Flags().BoolP("validate", "C", false, "Check the validity of the config file, then exit.")
	rootCmd.Flags().BoolP("verbose", "V", false, "Enables more verbose logging.")
	rootCmd.Flags().BoolP("version", "v", false, "Show the version number and quit.")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"config": carapace.ActionFiles(),
	})
}
