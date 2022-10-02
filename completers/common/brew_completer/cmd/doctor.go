package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var doctorCmd = &cobra.Command{
	Use:   "doctor",
	Short: "Check your system for potential problems",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(doctorCmd).Standalone()

	doctorCmd.Flags().BoolP("audit-debug", "D", false, "Enable debugging and profiling of audit methods.")
	doctorCmd.Flags().BoolP("debug", "d", false, "Display any debugging information.")
	doctorCmd.Flags().BoolP("help", "h", false, "Show this message.")
	doctorCmd.Flags().Bool("list-checks", false, "List all audit methods")
	doctorCmd.Flags().BoolP("quiet", "q", false, "Make some output more quiet.")
	doctorCmd.Flags().BoolP("verbose", "v", false, "Make some output more verbose.")
	rootCmd.AddCommand(doctorCmd)
}
