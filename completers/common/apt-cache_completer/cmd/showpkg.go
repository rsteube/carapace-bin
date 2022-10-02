package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/apt-cache_completer/cmd/action"
	"github.com/spf13/cobra"
)

var showpkgCmd = &cobra.Command{
	Use:   "showpkg",
	Short: "showpkg displays information about the packages",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(showpkgCmd).Standalone()

	rootCmd.AddCommand(showpkgCmd)

	carapace.Gen(showpkgCmd).PositionalAnyCompletion(
		action.ActionPackages(),
	)
}
