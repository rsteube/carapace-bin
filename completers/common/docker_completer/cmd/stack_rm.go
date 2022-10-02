package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var stack_rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove one or more stacks",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stack_rmCmd).Standalone()

	stack_rmCmd.Flags().String("orchestrator", "", "Orchestrator to use (swarm|kubernetes|all)")
	stackCmd.AddCommand(stack_rmCmd)

	carapace.Gen(stack_rmCmd).FlagCompletion(carapace.ActionMap{
		"orchestrator": carapace.ActionValues("swarm", "kubernetes", "all"),
	})

	carapace.Gen(stack_rmCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if stack_rmCmd.Flag("orchestrator").Changed {
				return docker.ActionStacks(stack_rmCmd.Flag("orchestrator").Value.String())
			} else {
				return docker.ActionStacks("all")
			}
		}),
	)
}
