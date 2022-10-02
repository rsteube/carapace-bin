package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/vercel_completer/cmd/action"
	"github.com/spf13/cobra"
)

var logsCmd = &cobra.Command{
	Use:   "logs",
	Short: "Displays the logs for a deployment",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logsCmd).Standalone()

	rootCmd.AddCommand(logsCmd)

	carapace.Gen(logsCmd).PositionalCompletion(
		action.ActionDeployments(logsCmd),
	)
}
