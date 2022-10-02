package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var context_useCmd = &cobra.Command{
	Use:   "use",
	Short: "Set the current docker context",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(context_useCmd).Standalone()

	contextCmd.AddCommand(context_useCmd)

	carapace.Gen(context_useCmd).PositionalCompletion(
		docker.ActionContexts(),
	)
}
