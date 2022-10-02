package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/cargo_completer/cmd/action"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/cargo"
	"github.com/spf13/cobra"
)

var fixCmd = &cobra.Command{
	Use:   "fix",
	Short: "Automatically fix lint warnings reported by rustc",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fixCmd).Standalone()

	fixCmd.Flags().StringS("Z", "Z", "", "Unstable (nightly-only) flags to Cargo, see 'cargo -Z help' for details")
	fixCmd.Flags().Bool("all", false, "Alias for --workspace (deprecated)")
	fixCmd.Flags().Bool("all-features", false, "Activate all available features")
	fixCmd.Flags().Bool("all-targets", false, "Fix all targets (default)")
	fixCmd.Flags().Bool("allow-dirty", false, "Fix code even if the working directory is dirty")
	fixCmd.Flags().Bool("allow-no-vcs", false, "Fix code even if a VCS was not detected")
	fixCmd.Flags().Bool("allow-staged", false, "Fix code even if the working directory has staged changes")
	fixCmd.Flags().String("bench", "", "Fix only the specified bench target")
	fixCmd.Flags().Bool("benches", false, "Fix all benches")
	fixCmd.Flags().String("bin", "", "Fix only the specified binary")
	fixCmd.Flags().Bool("bins", false, "Fix all binaries")
	fixCmd.Flags().Bool("broken-code", false, "Fix code even if it already has compiler errors")
	fixCmd.Flags().String("color", "", "Coloring: auto, always, never")
	fixCmd.Flags().String("config", "", "Override a configuration value (unstable)")
	fixCmd.Flags().Bool("edition", false, "Fix in preparation for the next edition")
	fixCmd.Flags().Bool("edition-idioms", false, "Fix warnings to migrate to the idioms of an edition")
	fixCmd.Flags().String("example", "", "Fix only the specified example")
	fixCmd.Flags().Bool("examples", false, "Fix all examples")
	fixCmd.Flags().String("exclude", "", "Exclude packages from the fixes")
	fixCmd.Flags().String("features", "", "Space or comma separated list of features to activate")
	fixCmd.Flags().Bool("frozen", false, "Require Cargo.lock and cache are up to date")
	fixCmd.Flags().BoolP("help", "h", false, "Prints help information")
	fixCmd.Flags().Bool("ignore-rust-version", false, "Ignore `rust-version` specification in packages")
	fixCmd.Flags().StringP("jobs", "j", "", "Number of parallel jobs, defaults to # of CPUs")
	fixCmd.Flags().Bool("lib", false, "Fix only this package's library")
	fixCmd.Flags().Bool("locked", false, "Require Cargo.lock is up to date")
	fixCmd.Flags().String("manifest-path", "", "Path to Cargo.toml")
	fixCmd.Flags().String("message-format", "", "Error format")
	fixCmd.Flags().Bool("no-default-features", false, "Do not activate the `default` feature")
	fixCmd.Flags().Bool("offline", false, "Run without accessing the network")
	fixCmd.Flags().StringP("package", "p", "", "Package(s) to fix")
	fixCmd.Flags().String("profile", "", "Build artifacts with the specified profile")
	fixCmd.Flags().BoolP("quiet", "q", false, "Do not print cargo log messages")
	fixCmd.Flags().Bool("release", false, "Fix artifacts in release mode, with optimizations")
	fixCmd.Flags().String("target", "", "Fix for the target triple")
	fixCmd.Flags().String("target-dir", "", "Directory for all generated artifacts")
	fixCmd.Flags().String("test", "", "Fix only the specified test target")
	fixCmd.Flags().Bool("tests", false, "Fix all tests")
	fixCmd.Flags().BoolP("verbose", "v", false, "Use verbose output (-vv very verbose/build.rs output)")
	fixCmd.Flags().Bool("workspace", false, "Fix all packages in the workspace")
	rootCmd.AddCommand(fixCmd)

	carapace.Gen(fixCmd).FlagCompletion(carapace.ActionMap{
		"benches": action.ActionTargets(fixCmd, action.TargetOpts{Bench: true}),
		"bin":     action.ActionTargets(fixCmd, action.TargetOpts{Bin: true}),
		"color":   action.ActionColorModes(),
		"example": action.ActionTargets(fixCmd, action.TargetOpts{Example: true}),
		"exclude": action.ActionWorkspaceMembers(fixCmd),
		"features": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return cargo.ActionFeatures(fixCmd.Flag("manifest-path").Value.String()).Invoke(c).Filter(c.Parts).ToA()
		}),
		"manifest-path":  carapace.ActionFiles(),
		"message-format": action.ActionMessageFormats(),
		"package":        action.ActionDependencies(fixCmd, true),
		"profile":        action.ActionProfiles(fixCmd),
		"target-dir":     carapace.ActionDirectories(),
		"test":           action.ActionTargets(fixCmd, action.TargetOpts{Test: true}),
	})
}
