package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var context_inspectCmd = &cobra.Command{
	Use:   "inspect",
	Short: "Display detailed information on one or more contexts",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(context_inspectCmd).Standalone()

	context_inspectCmd.Flags().StringP("format", "f", "", "Format the output using the given Go template")
	contextCmd.AddCommand(context_inspectCmd)

	carapace.Gen(context_inspectCmd).PositionalAnyCompletion(docker.ActionContexts())
}
