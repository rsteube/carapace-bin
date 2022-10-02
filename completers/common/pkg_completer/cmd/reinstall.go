package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/pkg_completer/cmd/action"
	"github.com/spf13/cobra"
)

var reinstallCmd = &cobra.Command{
	Use:   "reinstall",
	Short: "Reinstall specified installed packages at the latest version",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(reinstallCmd).Standalone()

	carapace.Gen(reinstallCmd).PositionalAnyCompletion(
		action.ActionPackages(),
	)

	rootCmd.AddCommand(reinstallCmd)
}
