package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/cargo_completer/cmd/action"
	"github.com/spf13/cobra"
)

var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fetchCmd).Standalone()

	fetchCmd.Flags().StringS("Z", "Z", "", "Unstable (nightly-only) flags to Cargo, see 'cargo -Z help' for details")
	fetchCmd.Flags().String("color", "", "Coloring: auto, always, never")
	fetchCmd.Flags().Bool("frozen", false, "Require Cargo.lock and cache are up to date")
	fetchCmd.Flags().BoolP("help", "h", false, "Prints help information")
	fetchCmd.Flags().Bool("locked", false, "Require Cargo.lock is up to date")
	fetchCmd.Flags().String("manifest-path", "", "Path to Cargo.toml")
	fetchCmd.Flags().Bool("offline", false, "Run without accessing the network")
	fetchCmd.Flags().BoolP("quiet", "q", false, "No output printed to stdout")
	fetchCmd.Flags().String("target", "", "Fetch dependencies for the target triple")
	fetchCmd.Flags().BoolP("verbose", "v", false, "Use verbose output (-vv very verbose/build.rs output)")
	rootCmd.AddCommand(fetchCmd)

	carapace.Gen(fetchCmd).FlagCompletion(carapace.ActionMap{
		"color":         action.ActionColorModes(),
		"manifest-path": carapace.ActionFiles(),
	})
}
