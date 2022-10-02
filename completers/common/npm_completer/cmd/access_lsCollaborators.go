package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/npm_completer/cmd/action"
	"github.com/spf13/cobra"
)

var access_lsCollaboratorsCmd = &cobra.Command{
	Use:   "ls-collaborators",
	Short: "list collaborators",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(access_lsCollaboratorsCmd).Standalone()
	accessCmd.AddCommand(access_lsCollaboratorsCmd)

	carapace.Gen(access_lsCollaboratorsCmd).PositionalCompletion(
		action.ActionPackages(access_lsCollaboratorsCmd),
		carapace.ActionValues(), // TODO <user>
	)
}
