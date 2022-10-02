package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var context_rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove one or more contexts",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(context_rmCmd).Standalone()

	context_rmCmd.Flags().BoolP("force", "f", false, "Force the removal of a context in use")
	contextCmd.AddCommand(context_rmCmd)

	carapace.Gen(context_rmCmd).PositionalAnyCompletion(docker.ActionContexts())
}
