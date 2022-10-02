package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var hybridSleepCmd = &cobra.Command{
	Use:   "hybrid-sleep",
	Short: "Hibernate and suspend the system",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(hybridSleepCmd).Standalone()

	rootCmd.AddCommand(hybridSleepCmd)
}
