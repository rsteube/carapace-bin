package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/cargo_completer/cmd/action"
	"github.com/spf13/cobra"
)

var benchCmd = &cobra.Command{
	Use:     "bench",
	Short:   "Execute all benchmarks of a local package",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groupFor("bench"),
}

func init() {
	carapace.Gen(benchCmd).Standalone()

	benchCmd.Flags().Bool("all", false, "Alias for --workspace (deprecated)")
	benchCmd.Flags().Bool("all-features", false, "Activate all available features")
	benchCmd.Flags().Bool("all-targets", false, "Benchmark all targets")
	benchCmd.Flags().StringSlice("bench", []string{}, "Benchmark only the specified bench target")
	benchCmd.Flags().Bool("benches", false, "Benchmark all benches")
	benchCmd.Flags().StringSlice("bin", []string{}, "Benchmark only the specified binary")
	benchCmd.Flags().Bool("bins", false, "Benchmark all binaries")
	benchCmd.Flags().StringSlice("example", []string{}, "Benchmark only the specified example")
	benchCmd.Flags().Bool("examples", false, "Benchmark all examples")
	benchCmd.Flags().StringSlice("exclude", []string{}, "Exclude packages from the benchmark")
	benchCmd.Flags().StringSliceP("features", "F", []string{}, "Space or comma separated list of features to activate")
	benchCmd.Flags().BoolP("help", "h", false, "Print help")
	benchCmd.Flags().Bool("ignore-rust-version", false, "Ignore `rust-version` specification in packages")
	benchCmd.Flags().StringP("jobs", "j", "", "Number of parallel jobs, defaults to # of CPUs")
	benchCmd.Flags().Bool("keep-going", false, "Do not abort the build as soon as there is an error (unstable)")
	benchCmd.Flags().Bool("lib", false, "Benchmark only this package's library")
	benchCmd.Flags().String("manifest-path", "", "Path to Cargo.toml")
	benchCmd.Flags().StringSlice("message-format", []string{}, "Error format")
	benchCmd.Flags().Bool("no-default-features", false, "Do not activate the `default` feature")
	benchCmd.Flags().Bool("no-fail-fast", false, "Run all benchmarks regardless of failure")
	benchCmd.Flags().Bool("no-run", false, "Compile, but don't run benchmarks")
	benchCmd.Flags().StringSliceP("package", "p", []string{}, "Package to run benchmarks for")
	benchCmd.Flags().String("profile", "", "Build artifacts with the specified profile")
	benchCmd.Flags().BoolP("quiet", "q", false, "Do not print cargo log messages")
	benchCmd.Flags().StringSlice("target", []string{}, "Build for the target triple")
	benchCmd.Flags().String("target-dir", "", "Directory for all generated artifacts")
	benchCmd.Flags().StringSlice("test", []string{}, "Benchmark only the specified test target")
	benchCmd.Flags().Bool("tests", false, "Benchmark all tests")
	benchCmd.Flags().String("timings", "", "Timing output formats (unstable) (comma separated): html, json")
	benchCmd.Flags().Bool("unit-graph", false, "Output build graph in JSON (unstable)")
	benchCmd.Flags().Bool("workspace", false, "Benchmark all packages in the workspace")
	benchCmd.Flag("timings").NoOptDefVal = " "
	rootCmd.AddCommand(benchCmd)

	carapace.Gen(benchCmd).FlagCompletion(carapace.ActionMap{
		"bench":          action.ActionTargets(benchCmd, action.TargetOpts{Bench: true}),
		"bin":            action.ActionTargets(benchCmd, action.TargetOpts{Bin: true}),
		"example":        action.ActionTargets(benchCmd, action.TargetOpts{Example: true}),
		"exclude":        action.ActionWorkspaceMembers(benchCmd),
		"features":       action.ActionFeatures(benchCmd).UniqueList(","),
		"manifest-path":  carapace.ActionFiles(),
		"message-format": action.ActionMessageFormats(),
		"package":        action.ActionDependencies(benchCmd, true),
		"profile":        action.ActionProfiles(benchCmd),
		"target-dir":     carapace.ActionDirectories(),
		"test":           action.ActionTargets(benchCmd, action.TargetOpts{Test: true}),
	})
}
