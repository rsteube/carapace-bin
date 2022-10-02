package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var runner_tokenCmd = &cobra.Command{
	Use:   "token",
	Short: "Operate on runner tokens",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(runner_tokenCmd).Standalone()
	runnerCmd.AddCommand(runner_tokenCmd)
}
