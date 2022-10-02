package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/cargo_completer/cmd/action"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update dependencies as recorded in the local lock file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(updateCmd).Standalone()

	updateCmd.Flags().StringS("Z", "Z", "", "Unstable (nightly-only) flags to Cargo, see 'cargo -Z help' for details")
	updateCmd.Flags().Bool("aggressive", false, "Force updating all dependencies of <name> as well")
	updateCmd.Flags().String("color", "", "Coloring: auto, always, never")
	updateCmd.Flags().Bool("dry-run", false, "Don't actually write the lockfile")
	updateCmd.Flags().Bool("frozen", false, "Require Cargo.lock and cache are up to date")
	updateCmd.Flags().BoolP("help", "h", false, "Prints help information")
	updateCmd.Flags().Bool("locked", false, "Require Cargo.lock is up to date")
	updateCmd.Flags().String("manifest-path", "", "Path to Cargo.toml")
	updateCmd.Flags().Bool("offline", false, "Run without accessing the network")
	updateCmd.Flags().StringP("package", "p", "", "Package to update")
	updateCmd.Flags().String("precise", "", "Update a single dependency to exactly PRECISE")
	updateCmd.Flags().BoolP("quiet", "q", false, "No output printed to stdout")
	updateCmd.Flags().BoolP("verbose", "v", false, "Use verbose output (-vv very verbose/build.rs output)")
	rootCmd.AddCommand(updateCmd)

	carapace.Gen(updateCmd).FlagCompletion(carapace.ActionMap{
		"color":         action.ActionColorModes(),
		"manifest-path": carapace.ActionFiles(),
		"package":       action.ActionDependencies(updateCmd, false),
	})
}
