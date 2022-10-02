package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/vercel_completer/cmd/action"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var env_rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove an Environment Variable",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(env_rmCmd).Standalone()

	envCmd.AddCommand(env_rmCmd)

	carapace.Gen(env_rmCmd).PositionalCompletion(
		action.ActionEnvironmentVariables(env_rmCmd),
		action.ActionEnvironments(),
		git.ActionRefs(git.RefOption{LocalBranches: true}),
	)
}
