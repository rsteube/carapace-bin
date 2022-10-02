package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var codespace_editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit a codespace",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codespace_editCmd).Standalone()
	codespace_editCmd.Flags().StringP("codespace", "c", "", "Name of the codespace")
	codespace_editCmd.Flags().StringP("display-name", "d", "", "Set the display name")
	codespace_editCmd.Flags().StringP("machine", "m", "", "Set hardware specifications for the VM")
	codespaceCmd.AddCommand(codespace_editCmd)

	carapace.Gen(codespace_editCmd).FlagCompletion(carapace.ActionMap{
		"codespace": action.ActionCodespaces(),
		"machine":   action.ActionCodespaceMachines(),
	})
}
