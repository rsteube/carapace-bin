package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/npm_completer/cmd/action"
	"github.com/spf13/cobra"
)

var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit an installed package",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(editCmd).Standalone()
	editCmd.Flags().String("editor", "", "editor to use")

	rootCmd.AddCommand(editCmd)

	carapace.Gen(editCmd).FlagCompletion(carapace.ActionMap{
		"editor": carapace.ActionFiles(),
	})

	carapace.Gen(editCmd).PositionalCompletion(
		action.ActionModules(),
	)
}
