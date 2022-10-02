package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/ng_completer/cmd/action"
	"github.com/spf13/cobra"
)

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Compiles an Angular app into an output directory",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(buildCmd).Standalone()

	buildCmd.Flags().String("allowed-common-js-dependencies", "", "A list of CommonJS packages that are allowed to be used without a build time warning.")
	buildCmd.Flags().Bool("aot", false, "Build using Ahead of Time compilation.")
	buildCmd.Flags().String("base-href", "", "Base url for the application being built.")
	buildCmd.Flags().Bool("build-optimizer", false, "Enables '@angular-devkit/build-optimizer' optimizations when using the 'aot' option.")
	buildCmd.Flags().Bool("common-chunk", false, "Generate a seperate bundle containing code used across multiple bundles.")
	buildCmd.Flags().StringP("configuration", "c", "", "One or more named builder configurations")
	buildCmd.Flags().String("cross-origin", "", "Define the crossorigin attribute setting of elements that provide CORS support.")
	buildCmd.Flags().Bool("delete-output-path", false, "Delete the output path before building.")
	buildCmd.Flags().Bool("extract-licenses", false, "Extract all licenses in a separate file.")
	buildCmd.Flags().String("i18n-missing-translation", "", "How to handle missing translations for i18n.")
	buildCmd.Flags().String("index", "", "Configures the generation of the application's HTML index.")
	buildCmd.Flags().String("inline-style-language", "", "The stylesheet language to use for the application's inline component styles.")
	buildCmd.Flags().Bool("localize", false, "Translate the bundles in one or more locales.")
	buildCmd.Flags().Bool("main", false, "The full path for the main entry point to the app, relative to the current workspace.")
	buildCmd.Flags().Bool("named-chunks", false, "Use file name for lazy loaded chunks.")
	buildCmd.Flags().String("ngsw-config-path", "", "Path to ngsw-config.json.")
	buildCmd.Flags().Bool("optimization", false, "Enables optimization of the build output.")
	buildCmd.Flags().String("output-hashing", "", "Define the output filename cache-busting hashing mode.")
	buildCmd.Flags().String("output-path", "", "The full path for the new output directory, relative to the current workspace.")
	buildCmd.Flags().String("poll", "", "Enable and define the file watching poll time period in milliseconds.")
	buildCmd.Flags().String("polyfills", "", "The full path for the polyfills file, relative to the current workspace.")
	buildCmd.Flags().Bool("preserve-symlinks", false, "Do not use the real path when resolving modules.")
	buildCmd.Flags().Bool("progress", false, "Log progress to the console while building.")
	buildCmd.Flags().String("resources-output-path", "", "The path where style resources will be placed, relative to outputPath.")
	buildCmd.Flags().Bool("service-worker", false, "Generates a service worker config for production builds.")
	buildCmd.Flags().Bool("source-map", false, "Output source maps for scripts and styles.")
	buildCmd.Flags().Bool("stats-json", false, "Generates a 'stats.json' file which can be analyzed using tools such as 'webpack-bundle-analyzer'.")
	buildCmd.Flags().Bool("subresource-integrity", false, "Enables the use of subresource integrity validation.")
	buildCmd.Flags().Bool("ts-config", false, "The full path for the TypeScript configuration file, relative to the current workspace.")
	buildCmd.Flags().Bool("vendor-chunk", false, "Generate a seperate bundle containing only vendor libraries.")
	buildCmd.Flags().Bool("verbose", false, "Adds more details to output logging.")
	buildCmd.Flags().Bool("watch", false, "Run build when files change.")
	buildCmd.Flags().String("web-worker-ts-config", "", "TypeScript configuration for Web Worker modules.")
	rootCmd.AddCommand(buildCmd)

	// TODO commonjs packages, named configurations
	carapace.Gen(buildCmd).FlagCompletion(carapace.ActionMap{
		"configuration": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			if len(c.Args) < 1 {
				return carapace.ActionValues()
			}
			return action.ActionBuilderConfigurations(c.Args[0], "build").Invoke(c).Filter(c.Parts).ToA()
		}),
		"cross-origin":             carapace.ActionValues("none", "anonymous", "use-credentials"),
		"i18n-missing-translation": carapace.ActionValues("warning", "error", "ignore"),
		"inline-style-language":    carapace.ActionValues("css", "less", "sass", "scss"),
		"ngsw-config-path":         carapace.ActionFiles(".json"),
		"output-hashing":           carapace.ActionValues("none", "all", "media", "bundle"),
		"output-path":              carapace.ActionDirectories(),
		"polyfills":                carapace.ActionFiles(),
		"resources-output-path":    carapace.ActionDirectories(),
		"web-worker-ts-config":     carapace.ActionFiles(),
	})

	carapace.Gen(buildCmd).PositionalCompletion(
		action.ActionProjects(),
	)
}
