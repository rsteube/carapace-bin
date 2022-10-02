package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/systemctl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var logLevelCmd = &cobra.Command{
	Use:   "log-level",
	Short: "Get/set logging threshold for manager",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logLevelCmd).Standalone()

	rootCmd.AddCommand(logLevelCmd)

	carapace.Gen(logLevelCmd).PositionalCompletion(
		action.ActionLogLevels(),
	)
}
