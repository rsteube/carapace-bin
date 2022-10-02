package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var codespace_jupyterCmd = &cobra.Command{
	Use:   "jupyter",
	Short: "Open a codespace in JupyterLab",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codespace_jupyterCmd).Standalone()
	codespace_jupyterCmd.Flags().StringP("codespace", "c", "", "Name of the codespace")
	codespaceCmd.AddCommand(codespace_jupyterCmd)

	carapace.Gen(codespace_jupyterCmd).FlagCompletion(carapace.ActionMap{
		"codespace": action.ActionCodespaces(),
	})
}
