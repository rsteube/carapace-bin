package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/cargo_completer/cmd/action"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/cargo"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install a Rust binary",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(installCmd).Standalone()

	installCmd.Flags().StringS("Z", "Z", "", "Unstable (nightly-only) flags to Cargo, see 'cargo -Z help' for details")
	installCmd.Flags().Bool("all-features", false, "Activate all available features")
	installCmd.Flags().String("bin", "", "Install only the specified binary")
	installCmd.Flags().Bool("bins", false, "Install all binaries")
	installCmd.Flags().String("branch", "", "Branch to use when installing from git")
	installCmd.Flags().String("color", "", "Coloring: auto, always, never")
	installCmd.Flags().Bool("debug", false, "Build in debug mode instead of release mode")
	installCmd.Flags().String("example", "", "Install only the specified example")
	installCmd.Flags().Bool("examples", false, "Install all examples")
	installCmd.Flags().String("features", "", "Space or comma separated list of features to activate")
	installCmd.Flags().BoolP("force", "f", false, "Force overwriting existing crates or binaries")
	installCmd.Flags().Bool("frozen", false, "Require Cargo.lock and cache are up to date")
	installCmd.Flags().String("git", "", "Git URL to install the specified crate from")
	installCmd.Flags().BoolP("help", "h", false, "Prints help information")
	installCmd.Flags().String("index", "", "Registry index to install from")
	installCmd.Flags().StringP("jobs", "j", "", "Number of parallel jobs, defaults to # of CPUs")
	installCmd.Flags().Bool("list", false, "list all installed packages and their versions")
	installCmd.Flags().Bool("locked", false, "Require Cargo.lock is up to date")
	installCmd.Flags().Bool("no-default-features", false, "Do not activate the `default` feature")
	installCmd.Flags().Bool("no-track", false, "Do not save tracking information")
	installCmd.Flags().Bool("offline", false, "Run without accessing the network")
	installCmd.Flags().String("path", "", "Filesystem path to local crate to install")
	installCmd.Flags().String("profile", "", "Install artifacts with the specified profile")
	installCmd.Flags().BoolP("quiet", "q", false, "No output printed to stdout")
	installCmd.Flags().String("registry", "", "Registry to use")
	installCmd.Flags().String("rev", "", "Specific commit to use when installing from git")
	installCmd.Flags().String("root", "", "Directory to install packages into")
	installCmd.Flags().String("tag", "", "Tag to use when installing from git")
	installCmd.Flags().String("target", "", "Build for the target triple")
	installCmd.Flags().String("target-dir", "", "Directory for all generated artifacts")
	installCmd.Flags().BoolP("verbose", "v", false, "Use verbose output (-vv very verbose/build.rs output)")
	installCmd.Flags().String("version", "", "Specify a version to install")
	rootCmd.AddCommand(installCmd)

	carapace.Gen(installCmd).FlagCompletion(carapace.ActionMap{
		"bin": action.ActionTargets(installCmd, action.TargetOpts{Bin: true}),
		// TODO "branch"
		"color":   action.ActionColorModes(),
		"example": action.ActionTargets(installCmd, action.TargetOpts{Example: true}),
		"features": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return cargo.ActionFeatures("").Invoke(c).Filter(c.Parts).ToA()
		}),
		"path":     carapace.ActionFiles(),
		"profile":  action.ActionProfiles(installCmd),
		"registry": action.ActionRegistries(),
		// TODO rev
		"root": carapace.ActionDirectories(),
		// TODO tag
		"target-dir": carapace.ActionDirectories(),
		// TODO version
	})
}
