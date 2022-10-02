package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/ng_completer/cmd/action"
	"github.com/spf13/cobra"
)

var generate_guardCmd = &cobra.Command{
	Use:   "guard",
	Short: "Generates a new, generic route guard definition in the given or default project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(generate_guardCmd).Standalone()

	generate_guardCmd.Flags().Bool("flat", false, "Creates the new files at the top level of the current project.")
	generate_guardCmd.Flags().String("implements", "", "Specifies which interfaces to implement.")
	generate_guardCmd.Flags().String("project", "", "The name of the project.")
	generate_guardCmd.Flags().Bool("skip-tests", false, "Do not create \"spec.ts\" test files for the new guard.")
	generateCmd.AddCommand(generate_guardCmd)

	carapace.Gen(generate_guardCmd).FlagCompletion(carapace.ActionMap{
		"project": action.ActionProjects(),
	})
}
