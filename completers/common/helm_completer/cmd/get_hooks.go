package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/helm_completer/cmd/action"
	"github.com/spf13/cobra"
)

var get_hooksCmd = &cobra.Command{
	Use:   "hooks",
	Short: "download all hooks for a named release",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	get_hooksCmd.Flags().Int("revision", 0, "get the named release with revision")
	getCmd.AddCommand(get_hooksCmd)

	carapace.Gen(get_hooksCmd).PositionalCompletion(
		action.ActionReleases(),
	)
}
