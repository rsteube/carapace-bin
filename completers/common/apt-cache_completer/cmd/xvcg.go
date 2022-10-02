package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/apt-cache_completer/cmd/action"
	"github.com/spf13/cobra"
)

var xvcgCmd = &cobra.Command{
	Use:   "xvcg",
	Short: "generates output suitable for use by xvcg",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(xvcgCmd).Standalone()

	rootCmd.AddCommand(xvcgCmd)

	carapace.Gen(xvcgCmd).PositionalAnyCompletion(
		action.ActionPackages(),
	)
}
