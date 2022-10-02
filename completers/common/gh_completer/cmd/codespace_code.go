package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var codespace_codeCmd = &cobra.Command{
	Use:   "code",
	Short: "Open a codespace in Visual Studio Code",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codespace_codeCmd).Standalone()
	codespace_codeCmd.Flags().StringP("codespace", "c", "", "Name of the codespace")
	codespace_codeCmd.Flags().Bool("insiders", false, "Use the insiders version of Visual Studio Code")
	codespace_codeCmd.Flags().BoolP("web", "w", false, "Use the web version of Visual Studio Code")
	codespaceCmd.AddCommand(codespace_codeCmd)

	carapace.Gen(codespace_codeCmd).FlagCompletion(carapace.ActionMap{
		"codespace": action.ActionCodespaces(),
	})
}
