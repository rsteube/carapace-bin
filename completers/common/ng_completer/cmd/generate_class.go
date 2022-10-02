package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/ng_completer/cmd/action"
	"github.com/spf13/cobra"
)

var generate_classCmd = &cobra.Command{
	Use:   "class",
	Short: "Creates a new, generic class definition in the given or default project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(generate_classCmd).Standalone()

	generate_classCmd.Flags().String("project", "", "The name of the project.")
	generate_classCmd.Flags().Bool("skip-tests", false, "Do not create \"spec.ts\" test files for the new class.")
	generate_classCmd.Flags().String("type", "", "Adds a developer-defined type to the filename, in the format \"name.type.ts\".")
	generateCmd.AddCommand(generate_classCmd)

	carapace.Gen(generate_classCmd).FlagCompletion(carapace.ActionMap{
		"project": action.ActionProjects(),
	})
}
