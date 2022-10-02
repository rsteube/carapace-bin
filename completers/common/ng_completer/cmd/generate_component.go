package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/ng_completer/cmd/action"
	"github.com/spf13/cobra"
)

var generate_componentCmd = &cobra.Command{
	Use:   "component",
	Short: "Creates a new, generic component definition in the given or default project.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(generate_componentCmd).Standalone()

	generate_componentCmd.Flags().StringP("change-detection", "c", "", "The change detection strategy to use in the new component.")
	generate_componentCmd.Flags().BoolP("display-block", "b", false, "Specifies if the style will contain :host { display: block; }.")
	generate_componentCmd.Flags().Bool("export", false, "The declaring NgModule exports this component.")
	generate_componentCmd.Flags().Bool("flat", false, "Create the new files at the top level of the current project.")
	generate_componentCmd.Flags().BoolP("inline-style", "s", false, "Include styles inline in the component.ts file.")
	generate_componentCmd.Flags().BoolP("inline-template", "t", false, "Include template inline in the component.ts file.")
	generate_componentCmd.Flags().StringP("module", "m", "", "The declaring NgModule.")
	generate_componentCmd.Flags().StringP("prefix", "p", "", "The prefix to apply to the generated component selector.")
	generate_componentCmd.Flags().String("project", "", "The name of the project.")
	generate_componentCmd.Flags().String("selector", "", "The HTML selector to use for this component.")
	generate_componentCmd.Flags().Bool("skip-import", false, "Do not import this component into the owning NgModule.")
	generate_componentCmd.Flags().Bool("skip-selector", false, "Specifies if the component should have a selector or not.")
	generate_componentCmd.Flags().Bool("skip-tests", false, "Do not create \"spec.ts\" test files for the new component.")
	generate_componentCmd.Flags().String("style", "", "The file extension or preprocessor to use for style files")
	generate_componentCmd.Flags().String("type", "", "Adds a developer-defined type to the filename, in the format \"name.type.ts\".")
	generate_componentCmd.Flags().StringP("view-encapsulation", "v", "", "The view encapsulation strategy to use in the new component.")
	generateCmd.AddCommand(generate_componentCmd)

	carapace.Gen(generate_componentCmd).FlagCompletion(carapace.ActionMap{
		"change-detection":   carapace.ActionValues("Default", "OnPush"),
		"project":            action.ActionProjects(),
		"style":              carapace.ActionValues("css", "scss", "sass", "less", "none"),
		"view-encapsulation": carapace.ActionValues("Emulated", "None", "ShadowDom"),
	})
}
