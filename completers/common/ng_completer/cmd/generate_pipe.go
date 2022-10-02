package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/ng_completer/cmd/action"
	"github.com/spf13/cobra"
)

var generate_pipeCmd = &cobra.Command{
	Use:   "pipe",
	Short: "Creates a new, generic pipe definition in the given or default project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(generate_pipeCmd).Standalone()

	generate_pipeCmd.Flags().Bool("export", false, "The declaring NgModule exports this pipe.")
	generate_pipeCmd.Flags().Bool("flat", false, "When true (the default) creates files at the top level of the project.")
	generate_pipeCmd.Flags().BoolP("module", "m", false, "The declaring NgModule.")
	generate_pipeCmd.Flags().String("project", "", "The name of the project.")
	generate_pipeCmd.Flags().Bool("skip-import", false, "Do not import this pipe into the owning NgModule.")
	generate_pipeCmd.Flags().Bool("skip-tests", false, "Do not create \"spec.ts\" test files for the new pipe.")
	generateCmd.AddCommand(generate_pipeCmd)

	carapace.Gen(generate_pipeCmd).FlagCompletion(carapace.ActionMap{
		"project": action.ActionProjects(),
	})
}
