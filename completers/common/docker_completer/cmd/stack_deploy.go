package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var stack_deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploy a new stack or update an existing stack",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stack_deployCmd).Standalone()

	stack_deployCmd.Flags().String("bundle-file", "", "Path to a Distributed Application Bundle file")
	stack_deployCmd.Flags().StringP("compose-file", "c", "", "Path to a Compose file, or \"-\" to read from stdin")
	stack_deployCmd.Flags().String("orchestrator", "", "Orchestrator to use (swarm|kubernetes|all)")
	stack_deployCmd.Flags().Bool("prune", false, "Prune services that are no longer referenced")
	stack_deployCmd.Flags().String("resolve-image", "", "Query the registry to resolve image digest and supported")
	stack_deployCmd.Flags().Bool("with-registry-auth", false, "Send registry authentication details to Swarm agents")
	stackCmd.AddCommand(stack_deployCmd)

	carapace.Gen(stack_deployCmd).FlagCompletion(carapace.ActionMap{
		"bundle-file":   carapace.ActionFiles(".dab"),
		"compose-file":  carapace.ActionFiles(".yml"),
		"orchestrator":  carapace.ActionValues("swarm", "kubernetes", "all"),
		"resolve-image": carapace.ActionValues("always", "changed", "never"),
	})

	carapace.Gen(stack_deployCmd).PositionalCompletion(
		docker.ActionStacks("all"),
	)
}
