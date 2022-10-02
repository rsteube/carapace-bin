package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var stack_servicesCmd = &cobra.Command{
	Use:   "services",
	Short: "List the services in the stack",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stack_servicesCmd).Standalone()

	stack_servicesCmd.Flags().StringP("filter", "f", "", "Filter output based on conditions provided")
	stack_servicesCmd.Flags().String("format", "", "Pretty-print services using a Go template")
	stack_servicesCmd.Flags().String("orchestrator", "", "Orchestrator to use (swarm|kubernetes|all)")
	stack_servicesCmd.Flags().BoolP("quiet", "q", false, "Only display IDs")
	stackCmd.AddCommand(stack_servicesCmd)

	carapace.Gen(stack_servicesCmd).FlagCompletion(carapace.ActionMap{
		"orchestrator": carapace.ActionValues("swarm", "kubernetes", "all"),
	})

	carapace.Gen(stack_servicesCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if stack_servicesCmd.Flag("orchestrator").Changed {
				return docker.ActionStacks(stack_servicesCmd.Flag("orchestrator").Value.String())
			} else {
				return docker.ActionStacks("all")
			}
		}),
	)
}
