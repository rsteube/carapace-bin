package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/cargo_completer/cmd/action"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/cargo"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:     "run",
	Aliases: []string{"r"},
	Short:   "Run a binary or example of the local package",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(runCmd).Standalone()

	runCmd.Flags().StringS("Z", "Z", "", "Unstable (nightly-only) flags to Cargo, see 'cargo -Z help' for details")
	runCmd.Flags().Bool("all-features", false, "Activate all available features")
	runCmd.Flags().String("bin", "", "Name of the bin target to run")
	runCmd.Flags().String("color", "", "Coloring: auto, always, never")
	runCmd.Flags().String("example", "", "Name of the example target to run")
	runCmd.Flags().String("features", "", "Space or comma separated list of features to activate")
	runCmd.Flags().Bool("frozen", false, "Require Cargo.lock and cache are up to date")
	runCmd.Flags().BoolP("help", "h", false, "Prints help information")
	runCmd.Flags().StringP("jobs", "j", "", "Number of parallel jobs, defaults to # of CPUs")
	runCmd.Flags().Bool("locked", false, "Require Cargo.lock is up to date")
	runCmd.Flags().String("manifest-path", "", "Path to Cargo.toml")
	runCmd.Flags().String("message-format", "", "Error format")
	runCmd.Flags().Bool("no-default-features", false, "Do not activate the `default` feature")
	runCmd.Flags().Bool("offline", false, "Run without accessing the network")
	runCmd.Flags().StringP("package", "p", "", "Package with the target to run")
	runCmd.Flags().String("profile", "", "Build artifacts with the specified profile")
	runCmd.Flags().BoolP("quiet", "q", false, "No output printed to stdout")
	runCmd.Flags().Bool("release", false, "Build artifacts in release mode, with optimizations")
	runCmd.Flags().String("target", "", "Build for the target triple")
	runCmd.Flags().String("target-dir", "", "Directory for all generated artifacts")
	runCmd.Flags().BoolP("verbose", "v", false, "Use verbose output (-vv very verbose/build.rs output)")
	rootCmd.AddCommand(runCmd)

	carapace.Gen(runCmd).FlagCompletion(carapace.ActionMap{
		"bin":     action.ActionTargets(runCmd, action.TargetOpts{Bin: true}),
		"color":   action.ActionColorModes(),
		"example": action.ActionTargets(runCmd, action.TargetOpts{Example: true}),
		"features": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return cargo.ActionFeatures(runCmd.Flag("manifest-path").Value.String()).Invoke(c).Filter(c.Parts).ToA()
		}),
		"manifest-path":  carapace.ActionFiles(),
		"message-format": action.ActionMessageFormats(),
		"package":        action.ActionDependencies(runCmd, true),
		"profile":        action.ActionProfiles(runCmd),
		"target-dir":     carapace.ActionDirectories(),
	})
}
