package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/ng_completer/cmd/action"
	"github.com/spf13/cobra"
)

var generate_directiveCmd = &cobra.Command{
	Use:   "directive",
	Short: "Creates a new, generic directive definition in the given or default project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(generate_directiveCmd).Standalone()

	generate_directiveCmd.Flags().Bool("export", false, "The declaring NgModule exports this directive.")
	generate_directiveCmd.Flags().Bool("flat", false, "Creates the new files at the top level of the current project.")
	generate_directiveCmd.Flags().StringP("module", "m", "", "The declaring NgModule.")
	generate_directiveCmd.Flags().StringP("prefix", "p", "", "A prefix to apply to generated selectors.")
	generate_directiveCmd.Flags().String("project", "", "The name of the project.")
	generate_directiveCmd.Flags().String("selector", "", "The HTML selector to use for this directive.")
	generate_directiveCmd.Flags().Bool("skip-import", false, "Do not import this directive into the owning NgModule.")
	generate_directiveCmd.Flags().Bool("skip-tests", false, "Do not create \"spec.ts\" test files for the new class.")
	generateCmd.AddCommand(generate_directiveCmd)

	carapace.Gen(generate_directiveCmd).FlagCompletion(carapace.ActionMap{
		"project": action.ActionProjects(),
	})
}
