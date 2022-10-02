package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/ng_completer/cmd/action"
	"github.com/spf13/cobra"
)

var generate_resolverCmd = &cobra.Command{
	Use:   "resolver",
	Short: "Generates a new, generic resolver definition in the given or default project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(generate_resolverCmd).Standalone()

	generate_resolverCmd.Flags().Bool("flat", false, "Creates the new files at the top level of the current project.")
	generate_resolverCmd.Flags().String("project", "", "The name of the project.")
	generate_resolverCmd.Flags().Bool("skip-tests", false, "Do not create \"spec.ts\" test files for the new resolver.")
	generateCmd.AddCommand(generate_resolverCmd)

	carapace.Gen(generate_resolverCmd).FlagCompletion(carapace.ActionMap{
		"project": action.ActionProjects(),
	})
}
