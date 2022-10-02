package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/npm_completer/cmd/action"
	"github.com/spf13/cobra"
)

var access_grantCmd = &cobra.Command{
	Use:   "grant",
	Short: "grant access",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(access_grantCmd).Standalone()
	accessCmd.AddCommand(access_grantCmd)

	carapace.Gen(access_grantCmd).PositionalCompletion(
		carapace.ActionValues("read-only", "read-write"),
		carapace.ActionValues(), // TODO <scope:team>
		action.ActionPackages(access_grantCmd),
	)
}
