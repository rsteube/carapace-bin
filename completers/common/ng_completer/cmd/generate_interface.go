package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/ng_completer/cmd/action"
	"github.com/spf13/cobra"
)

var generate_interfaceCmd = &cobra.Command{
	Use:   "interface",
	Short: "Creates a new, generic interface definition in the given or default project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(generate_interfaceCmd).Standalone()

	generate_interfaceCmd.Flags().String("prefix", "", "A prefix to apply to generated selectors.")
	generate_interfaceCmd.Flags().String("project", "", "The name of the project.")
	generateCmd.AddCommand(generate_interfaceCmd)

	carapace.Gen(generate_interfaceCmd).FlagCompletion(carapace.ActionMap{
		"project": action.ActionProjects(),
	})
}
