package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create an Angular workspace",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(newCmd).Standalone()

	newCmd.Flags().StringP("collection", "c", "", "A collection of schematics to use in generating the initial application.")
	newCmd.Flags().Bool("commit", false, "Initial git repository commit information.")
	newCmd.Flags().Bool("create-application", false, "Create a new initial application project in the 'src' folder of the new workspace.")
	newCmd.Flags().Bool("defaults", false, "Disable interactive input prompts for options with a default.")
	newCmd.Flags().String("directory", "", "The directory name to create the workspace in.")
	newCmd.Flags().BoolP("dry-run", "d", false, "Run through and reports activity without writing out results.")
	newCmd.Flags().BoolP("force", "f", false, "Force overwriting of existing files.")
	newCmd.Flags().BoolP("inline-style", "s", false, "Include styles inline in the component TS file.")
	newCmd.Flags().BoolP("inline-template", "t", false, "Include template inline in the component TS file.")
	newCmd.Flags().Bool("interactive", false, "Enable interactive input prompts.")
	newCmd.Flags().Bool("minimal", false, "Create a workspace without any testing frameworks.")
	newCmd.Flags().String("new-project-root", "", "The path where new projects will be created, relative to the new workspace root.")
	newCmd.Flags().String("package-manager", "", "The package manager used to install dependencies.")
	newCmd.Flags().StringP("prefix", "p", "", "The prefix to apply to generated selectors for the initial project.")
	newCmd.Flags().Bool("routing", false, "Generate a routing module for the initial project.")
	newCmd.Flags().BoolP("skip-git", "g", false, "Do not initialize a git repository.")
	newCmd.Flags().Bool("skip-install", false, "Do not install dependency packages.")
	newCmd.Flags().BoolP("skip-tests", "S", false, "Do not generate \"spec.ts\" test files for the new project.")
	newCmd.Flags().Bool("strict", false, "Creates a workspace with stricter type checking and stricter bundle budgets settings.")
	newCmd.Flags().String("style", "", "The file extension or preprocessor to use for style files.")
	newCmd.Flags().String("view-encapsulation", "", "The view encapsulation strategy to use in the initial project.")
	rootCmd.AddCommand(newCmd)

	carapace.Gen(newCmd).FlagCompletion(carapace.ActionMap{
		"directory":          carapace.ActionDirectories(),
		"package-manager":    carapace.ActionValues("npm", "yarn", "pnpm", "cnpm"),
		"style":              carapace.ActionValues("css", "scss", "sass", "less"),
		"view-encapsulation": carapace.ActionValues("Emulated", "None", "ShadowDom"),
	})
}
