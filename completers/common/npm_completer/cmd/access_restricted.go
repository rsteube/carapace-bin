package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/npm_completer/cmd/action"
	"github.com/spf13/cobra"
)

var access_restrictedCmd = &cobra.Command{
	Use:   "restricted",
	Short: "set restricted access",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(access_restrictedCmd).Standalone()
	accessCmd.AddCommand(access_restrictedCmd)

	carapace.Gen(access_restrictedCmd).PositionalCompletion(
		action.ActionPackages(access_restrictedCmd),
	)
}
