package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/npm_completer/cmd/action"
	"github.com/spf13/cobra"
)

var access_publicCmd = &cobra.Command{
	Use:   "public",
	Short: "set public access",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(access_publicCmd).Standalone()
	accessCmd.AddCommand(access_publicCmd)

	carapace.Gen(access_publicCmd).PositionalCompletion(
		action.ActionPackages(access_publicCmd),
	)
}
