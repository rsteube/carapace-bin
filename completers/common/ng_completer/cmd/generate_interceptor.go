package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/ng_completer/cmd/action"
	"github.com/spf13/cobra"
)

var generate_interceptorCmd = &cobra.Command{
	Use:   "interceptor",
	Short: "Creates a new, generic interceptor definition in the given or default project.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(generate_interceptorCmd).Standalone()

	generate_interceptorCmd.Flags().Bool("flat", false, "Creates files at the top level of the project.")
	generate_interceptorCmd.Flags().String("project", "", "The name of the project.")
	generate_interceptorCmd.Flags().Bool("skip-tests", false, "Do not create \"spec.ts\" test files for the new interceptor.")
	generateCmd.AddCommand(generate_interceptorCmd)

	carapace.Gen(generate_interceptorCmd).FlagCompletion(carapace.ActionMap{
		"project": action.ActionProjects(),
	})
}
