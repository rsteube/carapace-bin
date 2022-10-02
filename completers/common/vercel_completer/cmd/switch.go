package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/vercel_completer/cmd/action"
	"github.com/spf13/cobra"
)

var switchCmd = &cobra.Command{
	Use:   "switch",
	Short: "Switches between teams and your personal account",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(switchCmd).Standalone()

	rootCmd.AddCommand(switchCmd)

	carapace.Gen(switchCmd).PositionalCompletion(
		action.ActionTeams(switchCmd),
	)
}
