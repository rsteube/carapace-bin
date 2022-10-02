package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/waypoint_completer/cmd/action"
	"github.com/spf13/cobra"
)

var runner_profile_inspectCmd = &cobra.Command{
	Use:   "inspect",
	Short: "Show detailed information about a runner profile",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(runner_profile_inspectCmd).Standalone()

	runner_profileCmd.AddCommand(runner_profile_inspectCmd)

	carapace.Gen(runner_profile_inspectCmd).PositionalCompletion(
		action.ActionRunnerProfiles(),
	)
}
