package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var generate_applicationCmd = &cobra.Command{
	Use:   "application",
	Short: "Generates a new basic app definition in the projects subfolder of the workspace",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(generate_applicationCmd).Standalone()

	generate_applicationCmd.Flags().BoolP("inline-style", "s", false, "Include styles inline in the root component.ts file.")
	generate_applicationCmd.Flags().BoolP("inline-template", "t", false, "Include template inline in the root component.ts file.")
	generate_applicationCmd.Flags().Bool("minimal", false, "Create a bare-bones project without any testing frameworks.")
	generate_applicationCmd.Flags().BoolP("prefix", "p", false, "A prefix to apply to generated selectors.")
	generate_applicationCmd.Flags().Bool("routing", false, "Create a routing NgModule.")
	generate_applicationCmd.Flags().Bool("skip-install", false, "Skip installing dependency packages.")
	generate_applicationCmd.Flags().Bool("skip-package-json", false, "Do not add dependencies to the \"package.json\" file.")
	generate_applicationCmd.Flags().BoolP("skip-tests", "S", false, "Do not create \"spec.ts\" test files for the application.")
	generate_applicationCmd.Flags().Bool("strict", false, "Creates an application with stricter bundle budgets settings.")
	generate_applicationCmd.Flags().String("style", "", "The file extension or preprocessor to use for style files.")
	generate_applicationCmd.Flags().String("view-encapsulation", "", "The view encapsulation strategy to use in the new application.")
	generateCmd.AddCommand(generate_applicationCmd)

	carapace.Gen(generate_applicationCmd).FlagCompletion(carapace.ActionMap{
		"style":              carapace.ActionValues("css", "scss", "sass", "less"),
		"view-encapsulation": carapace.ActionValues("Emulated", "None", "ShadowDom"),
	})
}
