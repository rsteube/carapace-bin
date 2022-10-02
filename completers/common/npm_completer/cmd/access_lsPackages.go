package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var access_lsPackagesCmd = &cobra.Command{
	Use:   "ls-packages",
	Short: "list packages",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(access_lsPackagesCmd).Standalone()
	accessCmd.AddCommand(access_lsPackagesCmd)

	carapace.Gen(access_lsPackagesCmd).PositionalCompletion(
		carapace.ActionValues(), // TODO [<user>|<scope>|<scope:team>]
	)
}
