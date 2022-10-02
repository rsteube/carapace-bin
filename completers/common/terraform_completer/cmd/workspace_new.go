package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/terraform_completer/cmd/action"
	"github.com/spf13/cobra"
)

var workspace_newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a new workspace",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_newCmd).Standalone()

	workspace_newCmd.Flags().StringS("lock", "lock", "", "Don't hold a state lock during the operation.")
	workspace_newCmd.Flags().StringS("lock-timeout", "lock-timeout", "", "Duration to retry a state lock.")
	workspace_newCmd.Flags().StringS("state", "state", "", "Copy an existing state file into the new workspace.")
	workspaceCmd.AddCommand(workspace_newCmd)

	workspace_newCmd.Flag("lock").NoOptDefVal = " "
	workspace_newCmd.Flag("lock-timeout").NoOptDefVal = " "
	workspace_newCmd.Flag("state").NoOptDefVal = " "

	carapace.Gen(workspace_newCmd).FlagCompletion(carapace.ActionMap{
		"lock":  action.ActionBool(),
		"state": carapace.ActionFiles(),
	})

	carapace.Gen(workspace_newCmd).PositionalCompletion(
		action.ActionWorkspaces(),
	)
}
