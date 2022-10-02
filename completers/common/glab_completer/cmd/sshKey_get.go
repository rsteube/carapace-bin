package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var sshKey_getCmd = &cobra.Command{
	Use:   "get",
	Short: "Gets a single key",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sshKey_getCmd).Standalone()

	sshKeyCmd.AddCommand(sshKey_getCmd)

	carapace.Gen(sshKey_getCmd).PositionalCompletion(
		action.ActionSshKeyIds(sshKeyCmd),
	)
}
