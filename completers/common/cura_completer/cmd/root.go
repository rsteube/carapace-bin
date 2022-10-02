package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cura",
	Short: "Powerful, easy-to-use 3D printing software",
	Long:  "https://ultimaker.com/software/ultimaker-cura",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("debug", false, "Turn on the debug mode by setting this option.")
	rootCmd.Flags().Bool("external-backend", false, "Use an externally started backend instead of starting")
	rootCmd.Flags().Bool("headless", false, "Hides all GUI elements.")
	rootCmd.Flags().BoolP("help", "h", false, "show help")
	rootCmd.Flags().Bool("single-instance", false, "run a single instance only")
	rootCmd.Flags().Bool("version", false, "show program's version number and exit")

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
