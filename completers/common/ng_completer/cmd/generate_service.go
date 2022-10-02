package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/ng_completer/cmd/action"
	"github.com/spf13/cobra"
)

var generate_serviceCmd = &cobra.Command{
	Use:   "service",
	Short: "Creates a new, generic service definition in the given or default project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(generate_serviceCmd).Standalone()

	generate_serviceCmd.Flags().Bool("flat", false, "Creates files at the top level of the project.")
	generate_serviceCmd.Flags().String("project", "", "The name of the project.")
	generate_serviceCmd.Flags().Bool("skip-tests", false, "Do not create \"spec.ts\" test files for the new service.")
	generateCmd.AddCommand(generate_serviceCmd)

	carapace.Gen(generate_serviceCmd).FlagCompletion(carapace.ActionMap{
		"project": action.ActionProjects(),
	})
}
