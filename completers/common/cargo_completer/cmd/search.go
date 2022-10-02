package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/cargo_completer/cmd/action"
	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search packages in crates.io",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(searchCmd).Standalone()

	searchCmd.Flags().StringS("Z", "Z", "", "Unstable (nightly-only) flags to Cargo, see 'cargo -Z help' for details")
	searchCmd.Flags().String("color", "", "Coloring: auto, always, never")
	searchCmd.Flags().Bool("frozen", false, "Require Cargo.lock and cache are up to date")
	searchCmd.Flags().BoolP("help", "h", false, "Prints help information")
	searchCmd.Flags().String("index", "", "Registry index URL to upload the package to")
	searchCmd.Flags().String("limit", "", "Limit the number of results (default: 10, max: 100)")
	searchCmd.Flags().Bool("locked", false, "Require Cargo.lock is up to date")
	searchCmd.Flags().Bool("offline", false, "Run without accessing the network")
	searchCmd.Flags().BoolP("quiet", "q", false, "No output printed to stdout")
	searchCmd.Flags().String("registry", "", "Registry to use")
	searchCmd.Flags().BoolP("verbose", "v", false, "Use verbose output (-vv very verbose/build.rs output)")
	rootCmd.AddCommand(searchCmd)

	carapace.Gen(searchCmd).FlagCompletion(carapace.ActionMap{
		"color":    action.ActionColorModes(),
		"registry": action.ActionRegistries(),
	})
}
