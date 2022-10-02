package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/apt-cache_completer/cmd/action"
	"github.com/spf13/cobra"
)

var rdependsCmd = &cobra.Command{
	Use:   "rdepends",
	Short: "shows a listing of each dependency",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rdependsCmd).Standalone()

	rootCmd.AddCommand(rdependsCmd)

	carapace.Gen(rdependsCmd).PositionalAnyCompletion(
		action.ActionPackages(),
	)
}
