package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/circleci_completer/cmd/action"
	"github.com/spf13/cobra"
)

var context_storeSecretCmd = &cobra.Command{
	Use:   "store-secret",
	Short: "Store a new environment variable in the named context. The value is read from stdin.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(context_storeSecretCmd).Standalone()
	contextCmd.AddCommand(context_storeSecretCmd)

	// TODO org/context/secret
	carapace.Gen(context_storeSecretCmd).PositionalCompletion(
		action.ActionVcsTypes(),
	)
}
