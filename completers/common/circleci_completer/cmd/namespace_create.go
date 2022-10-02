package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/circleci_completer/cmd/action"
	"github.com/spf13/cobra"
)

var namespace_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a namespace",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(namespace_createCmd).Standalone()
	namespace_createCmd.Flags().Bool("integration-testing", false, "Enable test mode to bypass interactive UI.")
	namespace_createCmd.Flags().Bool("no-prompt", false, "Disable prompt to bypass interactive UI.")
	namespaceCmd.AddCommand(namespace_createCmd)

	// TODO org-name
	carapace.Gen(namespace_createCmd).PositionalCompletion(
		carapace.ActionValues(),
		action.ActionVcsTypes(),
	)
}
