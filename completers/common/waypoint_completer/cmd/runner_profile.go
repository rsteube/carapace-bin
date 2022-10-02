package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var runner_profileCmd = &cobra.Command{
	Use:   "profile",
	Short: "Manage runner profiles",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(runner_profileCmd).Standalone()

	runnerCmd.AddCommand(runner_profileCmd)
}
