package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/cargo_completer/cmd/action"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/cargo"
	"github.com/spf13/cobra"
)

var testCmd = &cobra.Command{
	Use:     "test",
	Aliases: []string{"t"},
	Short:   "Execute all unit and integration tests and build examples of a local package",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(testCmd).Standalone()

	testCmd.Flags().StringS("Z", "Z", "", "Unstable (nightly-only) flags to Cargo, see 'cargo -Z help' for details")
	testCmd.Flags().Bool("all", false, "Alias for --workspace (deprecated)")
	testCmd.Flags().Bool("all-features", false, "Activate all available features")
	testCmd.Flags().Bool("all-targets", false, "Test all targets")
	testCmd.Flags().String("bench", "", "Test only the specified bench target")
	testCmd.Flags().Bool("benches", false, "Test all benches")
	testCmd.Flags().String("bin", "", "Test only the specified binary")
	testCmd.Flags().Bool("bins", false, "Test all binaries")
	testCmd.Flags().String("color", "", "Coloring: auto, always, never")
	testCmd.Flags().Bool("doc", false, "Test only this library's documentation")
	testCmd.Flags().String("example", "", "Test only the specified example")
	testCmd.Flags().Bool("examples", false, "Test all examples")
	testCmd.Flags().String("exclude", "", "Exclude packages from the test")
	testCmd.Flags().String("features", "", "Space or comma separated list of features to activate")
	testCmd.Flags().Bool("frozen", false, "Require Cargo.lock and cache are up to date")
	testCmd.Flags().BoolP("help", "h", false, "Prints help information")
	testCmd.Flags().StringP("jobs", "j", "", "Number of parallel jobs, defaults to # of CPUs")
	testCmd.Flags().Bool("lib", false, "Test only this package's library unit tests")
	testCmd.Flags().Bool("locked", false, "Require Cargo.lock is up to date")
	testCmd.Flags().String("manifest-path", "", "Path to Cargo.toml")
	testCmd.Flags().String("message-format", "", "Error format")
	testCmd.Flags().Bool("no-default-features", false, "Do not activate the `default` feature")
	testCmd.Flags().Bool("no-fail-fast", false, "Run all tests regardless of failure")
	testCmd.Flags().Bool("no-run", false, "Compile, but don't run tests")
	testCmd.Flags().Bool("offline", false, "Run without accessing the network")
	testCmd.Flags().StringP("package", "p", "", "Package to run tests for")
	testCmd.Flags().String("profile", "", "Build artifacts with the specified profile")
	testCmd.Flags().BoolP("quiet", "q", false, "Display one character per test instead of one line")
	testCmd.Flags().Bool("release", false, "Build artifacts in release mode, with optimizations")
	testCmd.Flags().String("target", "", "Build for the target triple")
	testCmd.Flags().String("target-dir", "", "Directory for all generated artifacts")
	testCmd.Flags().String("test", "", "Test only the specified test target")
	testCmd.Flags().Bool("tests", false, "Test all tests")
	testCmd.Flags().BoolP("verbose", "v", false, "Use verbose output (-vv very verbose/build.rs output)")
	testCmd.Flags().Bool("workspace", false, "Test all packages in the workspace")
	rootCmd.AddCommand(testCmd)

	carapace.Gen(testCmd).FlagCompletion(carapace.ActionMap{
		"bench":   action.ActionTargets(testCmd, action.TargetOpts{Bench: true}),
		"bin":     action.ActionTargets(testCmd, action.TargetOpts{Bin: true}),
		"color":   action.ActionColorModes(),
		"example": action.ActionTargets(testCmd, action.TargetOpts{Example: true}),
		"exclude": action.ActionWorkspaceMembers(testCmd),
		"features": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return cargo.ActionFeatures(testCmd.Flag("manifest-path").Value.String()).Invoke(c).Filter(c.Parts).ToA()
		}),
		"manifest-path":  carapace.ActionFiles(),
		"message-format": action.ActionMessageFormats(),
		"package":        action.ActionDependencies(testCmd, true),
		"profile":        action.ActionProfiles(testCmd),
		"target-dir":     carapace.ActionDirectories(),
		"test":           action.ActionTargets(testCmd, action.TargetOpts{Test: true}),
	})
}
