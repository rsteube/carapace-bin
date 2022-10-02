package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var doctorCmd = &cobra.Command{
	Use:   "doctor",
	Short: "Basic QMK environment checks",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(doctorCmd).Standalone()

	doctorCmd.Flags().BoolP("help", "h", false, "show this help message and exit")
	doctorCmd.Flags().BoolP("no", "n", false, "Answer no to all questions.")
	doctorCmd.Flags().BoolP("yes", "y", false, "Answer yes to all questions.")
	rootCmd.AddCommand(doctorCmd)
}
