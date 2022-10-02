package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var doctorCmd = &cobra.Command{
	Use:   "doctor",
	Short: "Show information about the installed tooling",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(doctorCmd).Standalone()

	doctorCmd.Flags().Bool("android-licenses", false, "Run the Android SDK manager tool to accept the SDK's licenses.")
	doctorCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	rootCmd.AddCommand(doctorCmd)
}
