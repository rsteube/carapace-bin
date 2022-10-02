package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/ng_completer/cmd/action"
	"github.com/spf13/cobra"
)

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Runs unit tests in a project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(testCmd).Standalone()

	testCmd.Flags().String("browsers", "", "Override which browsers tests are run against.")
	testCmd.Flags().Bool("code-coverage", false, "Output a code coverage report.")
	testCmd.Flags().String("code-coverage-exclude", "", "Globs to exclude from code coverage.")
	testCmd.Flags().StringP("configuration", "c", "", "One or more named builder configurations")
	testCmd.Flags().String("include", "", "Globs of files to include, relative to workspace or project root.")
	testCmd.Flags().String("inline-style-language", "", "The stylesheet language to use for the application's inline component styles.")
	testCmd.Flags().String("karma-config", "", "The name of the Karma configuration file.")
	testCmd.Flags().String("main", "", "The name of the main entry-point file.")
	testCmd.Flags().String("poll", "", "Enable and define the file watching poll time period in milliseconds.")
	testCmd.Flags().String("polyfills", "", "The name of the polyfills file.")
	testCmd.Flags().Bool("preserve-symlinks", false, "Do not use the real path when resolving modules.")
	testCmd.Flags().Bool("progress", false, "Log progress to the console while building.")
	testCmd.Flags().String("reporters", "", "Karma reporters to use. Directly passed to the karma runner.")
	testCmd.Flags().Bool("source-map", false, "Output source maps for scripts and styles.")
	testCmd.Flags().String("ts-config", "", "The name of the TypeScript configuration file.")
	testCmd.Flags().Bool("watch", false, "Run build when files change.")
	testCmd.Flags().String("web-worker-ts-config", "", "TypeScript configuration for Web Worker modules.")
	rootCmd.AddCommand(testCmd)

	carapace.Gen(testCmd).FlagCompletion(carapace.ActionMap{
		"configuration": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			if len(c.Args) < 1 {
				return carapace.ActionValues()
			}
			return action.ActionBuilderConfigurations(c.Args[0], "test").Invoke(c).Filter(c.Parts).ToA()
		}),
		"inline-style-language": carapace.ActionValues("css", "less", "sass", "scss"),
		"karma-config":          carapace.ActionFiles(),
		"main":                  carapace.ActionFiles(),
		"polyfills":             carapace.ActionFiles(),
		"ts-config":             carapace.ActionFiles(),
		"web-worker-ts-config":  carapace.ActionFiles(),
	})

	carapace.Gen(testCmd).PositionalCompletion(
		action.ActionProjects(),
	)
}
