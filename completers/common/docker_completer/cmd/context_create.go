package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var context_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a context",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(context_createCmd).Standalone()

	context_createCmd.Flags().String("default-stack-orchestrator", "", "Default orchestrator for stack operations to use with this context (swarm|kubernetes|all)")
	context_createCmd.Flags().String("description", "", "Description of the context")
	context_createCmd.Flags().String("docker", "", "set the docker endpoint (default [])")
	context_createCmd.Flags().String("from", "", "create context from a named context")
	context_createCmd.Flags().String("kubernetes", "", "set the kubernetes endpoint (default [])")
	contextCmd.AddCommand(context_createCmd)
}
