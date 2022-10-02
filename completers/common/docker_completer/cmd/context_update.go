package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var context_updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a context",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(context_updateCmd).Standalone()

	context_updateCmd.Flags().String("default-stack-orchestrator", "", "Default orchestrator for stack operations to use")
	context_updateCmd.Flags().String("description", "", "Description of the context")
	context_updateCmd.Flags().String("docker", "", "set the docker endpoint (default [])")
	context_updateCmd.Flags().String("kubernetes", "", "set the kubernetes endpoint (default [])")
	contextCmd.AddCommand(context_updateCmd)

	carapace.Gen(context_updateCmd).PositionalCompletion(
		docker.ActionContexts(),
	)
}
