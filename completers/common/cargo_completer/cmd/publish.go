package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/cargo_completer/cmd/action"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/cargo"
	"github.com/spf13/cobra"
)

var publishCmd = &cobra.Command{
	Use:   "publish",
	Short: "Upload a package to the registry",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(publishCmd).Standalone()

	publishCmd.Flags().StringS("Z", "Z", "", "Unstable (nightly-only) flags to Cargo, see 'cargo -Z help' for details")
	publishCmd.Flags().Bool("all-features", false, "Activate all available features")
	publishCmd.Flags().Bool("allow-dirty", false, "Allow dirty working directories to be packaged")
	publishCmd.Flags().String("color", "", "Coloring: auto, always, never")
	publishCmd.Flags().Bool("dry-run", false, "Perform all checks without uploading")
	publishCmd.Flags().String("features", "", "Space or comma separated list of features to activate")
	publishCmd.Flags().Bool("frozen", false, "Require Cargo.lock and cache are up to date")
	publishCmd.Flags().BoolP("help", "h", false, "Prints help information")
	publishCmd.Flags().String("index", "", "Registry index URL to upload the package to")
	publishCmd.Flags().StringP("jobs", "j", "", "Number of parallel jobs, defaults to # of CPUs")
	publishCmd.Flags().Bool("locked", false, "Require Cargo.lock is up to date")
	publishCmd.Flags().String("manifest-path", "", "Path to Cargo.toml")
	publishCmd.Flags().Bool("no-default-features", false, "Do not activate the `default` feature")
	publishCmd.Flags().Bool("no-verify", false, "Don't verify the contents by building them")
	publishCmd.Flags().Bool("offline", false, "Run without accessing the network")
	publishCmd.Flags().BoolP("quiet", "q", false, "No output printed to stdout")
	publishCmd.Flags().String("registry", "", "Registry to publish to")
	publishCmd.Flags().String("target", "", "Build for the target triple")
	publishCmd.Flags().String("target-dir", "", "Directory for all generated artifacts")
	publishCmd.Flags().String("token", "", "Token to use when uploading")
	publishCmd.Flags().BoolP("verbose", "v", false, "Use verbose output (-vv very verbose/build.rs output)")
	rootCmd.AddCommand(publishCmd)

	carapace.Gen(publishCmd).FlagCompletion(carapace.ActionMap{
		"color": action.ActionColorModes(),
		"features": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return cargo.ActionFeatures(publishCmd.Flag("manifest-path").Value.String()).Invoke(c).Filter(c.Parts).ToA()
		}),
		"manifest-path": carapace.ActionFiles(),
		"registry":      action.ActionRegistries(),
		"target-dir":    carapace.ActionDirectories(),
	})
}
