package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/npm_completer/cmd/action"
	"github.com/spf13/cobra"
)

var rebuildCmd = &cobra.Command{
	Use:   "rebuild",
	Short: "Rebuild a package",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rebuildCmd).Standalone()
	rebuildCmd.Flags().Bool("bin-links", false, "crete symlinks for package executables")
	rebuildCmd.Flags().BoolP("global", "g", false, "operate globally")
	rebuildCmd.Flags().Bool("ignore-scripts", false, "do not run scripts specified in package.json")
	addWorkspaceFlags(rebuildCmd)

	rootCmd.AddCommand(rebuildCmd)

	carapace.Gen(rebuildCmd).PositionalAnyCompletion(
		action.ActionPackages(rebuildCmd),
	)
}
