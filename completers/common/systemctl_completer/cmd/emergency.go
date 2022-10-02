package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var emergencyCmd = &cobra.Command{
	Use:   "emergency",
	Short: "Enter system emergency mode",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(emergencyCmd).Standalone()

	rootCmd.AddCommand(emergencyCmd)
}
