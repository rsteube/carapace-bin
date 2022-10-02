package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/cargo_completer/cmd/action"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/cargo"
	"github.com/spf13/cobra"
)

var docCmd = &cobra.Command{
	Use:   "doc",
	Short: "Build a package's documentation",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(docCmd).Standalone()

	docCmd.Flags().StringS("Z", "Z", "", "Unstable (nightly-only) flags to Cargo, see 'cargo -Z help' for details")
	docCmd.Flags().Bool("all", false, "Alias for --workspace (deprecated)")
	docCmd.Flags().Bool("all-features", false, "Activate all available features")
	docCmd.Flags().String("bin", "", "Document only the specified binary")
	docCmd.Flags().Bool("bins", false, "Document all binaries")
	docCmd.Flags().String("color", "", "Coloring: auto, always, never")
	docCmd.Flags().Bool("document-private-items", false, "Document private items")
	docCmd.Flags().String("exclude", "", "Exclude packages from the build")
	docCmd.Flags().String("features", "", "Space or comma separated list of features to activate")
	docCmd.Flags().Bool("frozen", false, "Require Cargo.lock and cache are up to date")
	docCmd.Flags().BoolP("help", "h", false, "Prints help information")
	docCmd.Flags().StringP("jobs", "j", "", "Number of parallel jobs, defaults to # of CPUs")
	docCmd.Flags().Bool("lib", false, "Document only this package's library")
	docCmd.Flags().Bool("locked", false, "Require Cargo.lock is up to date")
	docCmd.Flags().String("manifest-path", "", "Path to Cargo.toml")
	docCmd.Flags().String("message-format", "", "Error format")
	docCmd.Flags().Bool("no-default-features", false, "Do not activate the `default` feature")
	docCmd.Flags().Bool("no-deps", false, "Don't build documentation for dependencies")
	docCmd.Flags().Bool("offline", false, "Run without accessing the network")
	docCmd.Flags().Bool("open", false, "Opens the docs in a browser after the operation")
	docCmd.Flags().StringP("package", "p", "", "Package to document")
	docCmd.Flags().String("profile", "", "Build artifacts with the specified profile")
	docCmd.Flags().BoolP("quiet", "q", false, "No output printed to stdout")
	docCmd.Flags().Bool("release", false, "Build artifacts in release mode, with optimizations")
	docCmd.Flags().String("target", "", "Build for the target triple")
	docCmd.Flags().String("target-dir", "", "Directory for all generated artifacts")
	docCmd.Flags().BoolP("verbose", "v", false, "Use verbose output (-vv very verbose/build.rs output)")
	docCmd.Flags().Bool("workspace", false, "Document all packages in the workspace")
	rootCmd.AddCommand(docCmd)

	carapace.Gen(docCmd).FlagCompletion(carapace.ActionMap{
		"bin":     action.ActionTargets(docCmd, action.TargetOpts{Bin: true}),
		"color":   action.ActionColorModes(),
		"exclude": action.ActionWorkspaceMembers(docCmd),
		"features": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return cargo.ActionFeatures(docCmd.Flag("manifest-path").Value.String()).Invoke(c).Filter(c.Parts).ToA()
		}),
		"manifest-path":  carapace.ActionFiles(),
		"message-format": action.ActionMessageFormats(),
		"package":        action.ActionDependencies(docCmd, true),
		"profile":        action.ActionProfiles(docCmd),
		"target-dir":     carapace.ActionDirectories(),
	})
}
