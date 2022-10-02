package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/apt-cache_completer/cmd/action"
	"github.com/spf13/cobra"
)

var madisonCmd = &cobra.Command{
	Use:   "madison",
	Short: "mimic the output format of the Debian madison tool",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(madisonCmd).Standalone()

	rootCmd.AddCommand(madisonCmd)

	carapace.Gen(madisonCmd).PositionalAnyCompletion(
		action.ActionPackages(),
	)
}
