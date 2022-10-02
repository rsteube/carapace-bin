package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/apt-cache_completer/cmd/action"
	"github.com/spf13/cobra"
)

var showsrcCmd = &cobra.Command{
	Use:   "showsrc",
	Short: "showsrc displays all the source package records",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(showsrcCmd).Standalone()

	rootCmd.AddCommand(showsrcCmd)

	carapace.Gen(showsrcCmd).PositionalAnyCompletion(
		action.ActionPackages(),
	)
}
