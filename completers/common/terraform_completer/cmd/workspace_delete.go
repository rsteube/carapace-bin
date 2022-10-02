package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/terraform_completer/cmd/action"
	"github.com/spf13/cobra"
)

var workspace_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a workspace",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_deleteCmd).Standalone()

	workspace_deleteCmd.Flags().BoolS("force", "force", false, "Remove even a non-empty workspace.")
	workspace_deleteCmd.Flags().StringS("lock", "lock", "", "Don't hold a state lock during the operation.")
	workspace_deleteCmd.Flags().StringS("lock-timeout", "lock-timeout", "", "Duration to retry a state lock.")
	workspaceCmd.AddCommand(workspace_deleteCmd)

	workspace_deleteCmd.Flag("lock").NoOptDefVal = " "
	workspace_deleteCmd.Flag("lock-timeout").NoOptDefVal = " "

	carapace.Gen(workspace_deleteCmd).FlagCompletion(carapace.ActionMap{
		"lock": action.ActionBool(),
	})

	carapace.Gen(workspace_deleteCmd).PositionalCompletion(
		action.ActionWorkspaces(),
	)
}
