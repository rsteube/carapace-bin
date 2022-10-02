package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/npm_completer/cmd/action"
	"github.com/spf13/cobra"
)

var access_revokeCmd = &cobra.Command{
	Use:   "revoke",
	Short: "revoke access",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(access_revokeCmd).Standalone()
	accessCmd.AddCommand(access_revokeCmd)

	carapace.Gen(access_revokeCmd).PositionalCompletion(
		carapace.ActionValues(), // TODO <scope:team>
		action.ActionPackages(access_revokeCmd),
	)
}
