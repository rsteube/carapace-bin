package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var label_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a label from a repository",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(label_deleteCmd).Standalone()
	label_deleteCmd.Flags().Bool("confirm", false, "Confirm deletion without prompting")
	labelCmd.AddCommand(label_deleteCmd)

	carapace.Gen(label_deleteCmd).PositionalCompletion(
		action.ActionLabels(label_deleteCmd),
	)
}
