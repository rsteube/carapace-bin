package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/ng_completer/cmd/action"
	"github.com/spf13/cobra"
)

var generate_enumCmd = &cobra.Command{
	Use:   "enum",
	Short: "Generates a new, generic enum definition for the given or default project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(generate_enumCmd).Standalone()

	generate_enumCmd.Flags().String("project", "", "The name of the project in which to create the enum.")
	generate_enumCmd.Flags().String("type", "", "Adds a developer-defined type to the filename, in the format \"name.type.ts\".")
	generateCmd.AddCommand(generate_enumCmd)

	carapace.Gen(generate_enumCmd).FlagCompletion(carapace.ActionMap{
		"project": action.ActionProjects(),
	})
}
