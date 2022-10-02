package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/cargo_completer/cmd/action"
	"github.com/spf13/cobra"
)

var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Remove artifacts that cargo has generated in the past",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanCmd).Standalone()

	cleanCmd.Flags().StringS("Z", "Z", "", "Unstable (nightly-only) flags to Cargo, see 'cargo -Z help' for details")
	cleanCmd.Flags().String("color", "", "Coloring: auto, always, never")
	cleanCmd.Flags().Bool("doc", false, "Whether or not to clean just the documentation directory")
	cleanCmd.Flags().Bool("frozen", false, "Require Cargo.lock and cache are up to date")
	cleanCmd.Flags().BoolP("help", "h", false, "Prints help information")
	cleanCmd.Flags().Bool("locked", false, "Require Cargo.lock is up to date")
	cleanCmd.Flags().String("manifest-path", "", "Path to Cargo.toml")
	cleanCmd.Flags().Bool("offline", false, "Run without accessing the network")
	cleanCmd.Flags().StringP("package", "p", "", "Package to clean artifacts for")
	cleanCmd.Flags().String("profile", "", "Clean artifacts of the specified profile")
	cleanCmd.Flags().BoolP("quiet", "q", false, "No output printed to stdout")
	cleanCmd.Flags().Bool("release", false, "Whether or not to clean release artifacts")
	cleanCmd.Flags().String("target", "", "Target triple to clean output for")
	cleanCmd.Flags().String("target-dir", "", "Directory for all generated artifacts")
	cleanCmd.Flags().BoolP("verbose", "v", false, "Use verbose output (-vv very verbose/build.rs output)")
	rootCmd.AddCommand(cleanCmd)

	carapace.Gen(cleanCmd).FlagCompletion(carapace.ActionMap{
		"color":         action.ActionColorModes(),
		"manifest-path": carapace.ActionFiles(),
		"package":       action.ActionDependencies(cleanCmd, true),
		"profile":       action.ActionProfiles(cleanCmd),
		"target-dir":    carapace.ActionDirectories(),
	})
}
