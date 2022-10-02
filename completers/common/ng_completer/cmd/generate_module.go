package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/ng_completer/cmd/action"
	"github.com/spf13/cobra"
)

var generate_moduleCmd = &cobra.Command{
	Use:   "module",
	Short: "Creates a new, generic NgModule definition in the given or default project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(generate_moduleCmd).Standalone()

	generate_moduleCmd.Flags().Bool("flat", false, "Create the new files at the top level of the current project root.")
	generate_moduleCmd.Flags().StringP("module", "m", "", "The declaring NgModule.")
	generate_moduleCmd.Flags().String("project", "", "The name of the project.")
	generate_moduleCmd.Flags().String("route", "", "The route path for a lazy-loaded module.")
	generate_moduleCmd.Flags().Bool("routing", false, "Create a routing module.")
	generate_moduleCmd.Flags().String("routing-scope", "", "The scope for the new routing module.")
	generateCmd.AddCommand(generate_moduleCmd)

	carapace.Gen(generate_moduleCmd).FlagCompletion(carapace.ActionMap{
		"project":       action.ActionProjects(),
		"routing-scope": carapace.ActionValues("Child", "Root"),
	})
}
