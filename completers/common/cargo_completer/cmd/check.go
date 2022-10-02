package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/cargo_completer/cmd/action"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/cargo"
	"github.com/spf13/cobra"
)

var checkCmd = &cobra.Command{
	Use:     "check",
	Aliases: []string{"c"},
	Short:   "Check a local package and all of its dependencies for errors",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(checkCmd).Standalone()

	checkCmd.Flags().StringS("Z", "Z", "", "Unstable (nightly-only) flags to Cargo, see 'cargo -Z help' for details")
	checkCmd.Flags().Bool("all", false, "Alias for --workspace (deprecated)")
	checkCmd.Flags().Bool("all-features", false, "Activate all available features")
	checkCmd.Flags().Bool("all-targets", false, "Check all targets")
	checkCmd.Flags().String("bench", "", "Check only the specified bench target")
	checkCmd.Flags().Bool("benches", false, "Check all benches")
	checkCmd.Flags().String("bin", "", "Check only the specified binary")
	checkCmd.Flags().Bool("bins", false, "Check all binaries")
	checkCmd.Flags().String("color", "", "Coloring: auto, always, never")
	checkCmd.Flags().String("example", "", "Check only the specified example")
	checkCmd.Flags().Bool("examples", false, "Check all examples")
	checkCmd.Flags().String("exclude", "", "Exclude packages from the check")
	checkCmd.Flags().String("features", "", "Space or comma separated list of features to activate")
	checkCmd.Flags().Bool("frozen", false, "Require Cargo.lock and cache are up to date")
	checkCmd.Flags().BoolP("help", "h", false, "Prints help information")
	checkCmd.Flags().StringP("jobs", "j", "", "Number of parallel jobs, defaults to # of CPUs")
	checkCmd.Flags().Bool("lib", false, "Check only this package's library")
	checkCmd.Flags().Bool("locked", false, "Require Cargo.lock is up to date")
	checkCmd.Flags().String("manifest-path", "", "Path to Cargo.toml")
	checkCmd.Flags().String("message-format", "", "Error format")
	checkCmd.Flags().Bool("no-default-features", false, "Do not activate the `default` feature")
	checkCmd.Flags().Bool("offline", false, "Run without accessing the network")
	checkCmd.Flags().StringP("package", "p", "", "Package(s) to check")
	checkCmd.Flags().String("profile", "", "Check artifacts with the specified profile")
	checkCmd.Flags().BoolP("quiet", "q", false, "No output printed to stdout")
	checkCmd.Flags().Bool("release", false, "Check artifacts in release mode, with optimizations")
	checkCmd.Flags().String("target", "", "Check for the target triple")
	checkCmd.Flags().String("target-dir", "", "Directory for all generated artifacts")
	checkCmd.Flags().String("test", "", "Check only the specified test target")
	checkCmd.Flags().Bool("tests", false, "Check all tests")
	checkCmd.Flags().BoolP("verbose", "v", false, "Use verbose output (-vv very verbose/build.rs output)")
	checkCmd.Flags().Bool("workspace", false, "Check all packages in the workspace")
	rootCmd.AddCommand(checkCmd)

	carapace.Gen(checkCmd).FlagCompletion(carapace.ActionMap{
		"bench":   action.ActionTargets(checkCmd, action.TargetOpts{Bench: true}),
		"bin":     action.ActionTargets(checkCmd, action.TargetOpts{Bin: true}),
		"color":   action.ActionColorModes(),
		"example": action.ActionTargets(checkCmd, action.TargetOpts{Example: true}),
		"exclude": action.ActionWorkspaceMembers(checkCmd),
		"features": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return cargo.ActionFeatures(checkCmd.Flag("manifest-path").Value.String())
		}),
		"manifest-path":  carapace.ActionFiles(),
		"message-format": action.ActionMessageFormats(),
		"package":        action.ActionDependencies(buildCmd, true),
		"profile":        action.ActionProfiles(checkCmd),
		"target-dir":     carapace.ActionDirectories(),
		"test":           action.ActionTargets(checkCmd, action.TargetOpts{Test: true}),
	})
}
