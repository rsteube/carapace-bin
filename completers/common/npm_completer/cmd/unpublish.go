package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/npm_completer/cmd/action"
	"github.com/spf13/cobra"
)

var unpublishCmd = &cobra.Command{
	Use:   "unpublish",
	Short: "Remove a package from the registry",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(unpublishCmd).Standalone()
	unpublishCmd.Flags().Bool("dry-run", false, "only report changes")
	unpublishCmd.Flags().BoolP("force", "f", false, "remove various protections against unfortunate side effects")
	addWorkspaceFlags(unpublishCmd)

	rootCmd.AddCommand(unpublishCmd)

	carapace.Gen(unpublishCmd).PositionalCompletion(
		action.ActionPackages(unpublishCmd),
	)
}
