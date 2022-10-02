package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/cargo_completer/cmd/action"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/cargo"
	"github.com/spf13/cobra"
)

var benchCmd = &cobra.Command{
	Use:   "bench",
	Short: "Execute all benchmarks of a local package",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(benchCmd).Standalone()

	benchCmd.Flags().StringS("Z", "Z", "", "Unstable (nightly-only) flags to Cargo, see 'cargo -Z help' for details")
	benchCmd.Flags().Bool("all", false, "Alias for --workspace (deprecated)")
	benchCmd.Flags().Bool("all-features", false, "Activate all available features")
	benchCmd.Flags().Bool("all-targets", false, "Benchmark all targets")
	benchCmd.Flags().String("bench", "", "Benchmark only the specified bench target")
	benchCmd.Flags().Bool("benches", false, "Benchmark all benches")
	benchCmd.Flags().String("bin", "", "Benchmark only the specified binary")
	benchCmd.Flags().Bool("bins", false, "Benchmark all binaries")
	benchCmd.Flags().String("color", "", "Coloring: auto, always, never")
	benchCmd.Flags().String("example", "", "Benchmark only the specified example")
	benchCmd.Flags().Bool("examples", false, "Benchmark all examples")
	benchCmd.Flags().String("exclude", "", "Exclude packages from the benchmark")
	benchCmd.Flags().String("features", "", "Space or comma separated list of features to activate")
	benchCmd.Flags().Bool("frozen", false, "Require Cargo.lock and cache are up to date")
	benchCmd.Flags().BoolP("help", "h", false, "Prints help information")
	benchCmd.Flags().StringP("jobs", "j", "", "Number of parallel jobs, defaults to # of CPUs")
	benchCmd.Flags().Bool("lib", false, "Benchmark only this package's library")
	benchCmd.Flags().Bool("locked", false, "Require Cargo.lock is up to date")
	benchCmd.Flags().String("manifest-path", "", "Path to Cargo.toml")
	benchCmd.Flags().String("message-format", "", "Error format")
	benchCmd.Flags().Bool("no-default-features", false, "Do not activate the `default` feature")
	benchCmd.Flags().Bool("no-fail-fast", false, "Run all benchmarks regardless of failure")
	benchCmd.Flags().Bool("no-run", false, "Compile, but don't run benchmarks")
	benchCmd.Flags().Bool("offline", false, "Run without accessing the network")
	benchCmd.Flags().StringP("package", "p", "", "Package to run benchmarks for")
	benchCmd.Flags().BoolP("quiet", "q", false, "No output printed to stdout")
	benchCmd.Flags().String("target", "", "Build for the target triple")
	benchCmd.Flags().String("target-dir", "", "Directory for all generated artifacts")
	benchCmd.Flags().String("test", "", "Benchmark only the specified test target")
	benchCmd.Flags().Bool("tests", false, "Benchmark all tests")
	benchCmd.Flags().BoolP("verbose", "v", false, "Use verbose output (-vv very verbose/build.rs output)")
	benchCmd.Flags().Bool("workspace", false, "Benchmark all packages in the workspace")
	rootCmd.AddCommand(benchCmd)

	carapace.Gen(benchCmd).FlagCompletion(carapace.ActionMap{
		"bench":   action.ActionTargets(benchCmd, action.TargetOpts{Bench: true}),
		"bin":     action.ActionTargets(benchCmd, action.TargetOpts{Bin: true}),
		"color":   action.ActionColorModes(),
		"example": action.ActionTargets(benchCmd, action.TargetOpts{Example: true}),
		"exclude": action.ActionWorkspaceMembers(benchCmd),
		"features": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return cargo.ActionFeatures(benchCmd.Flag("manifest-path").Value.String()).Invoke(c).Filter(c.Parts).ToA()
		}),
		"manifest-path":  carapace.ActionFiles(),
		"message-format": action.ActionMessageFormats(),
		"package":        action.ActionDependencies(benchCmd, true),
		"target-dir":     carapace.ActionDirectories(),
		"test":           action.ActionTargets(benchCmd, action.TargetOpts{Test: true}),
	})
}
