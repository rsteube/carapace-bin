package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var runnerCmd = &cobra.Command{
	Use:   "runner",
	Short: "Runner management",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(runnerCmd).Standalone()

	rootCmd.AddCommand(runnerCmd)
}
