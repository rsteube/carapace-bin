package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var codespace_logsCmd = &cobra.Command{
	Use:   "logs",
	Short: "Access codespace logs",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codespace_logsCmd).Standalone()
	codespace_logsCmd.Flags().StringP("codespace", "c", "", "Name of the codespace")
	codespace_logsCmd.Flags().BoolP("follow", "f", false, "Tail and follow the logs")
	codespaceCmd.AddCommand(codespace_logsCmd)

	carapace.Gen(codespace_logsCmd).FlagCompletion(carapace.ActionMap{
		"codespace": action.ActionCodespaces(),
	})
}
