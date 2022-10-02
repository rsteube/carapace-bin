package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/npm_completer/cmd/action"
	"github.com/spf13/cobra"
)

var pruneCmd = &cobra.Command{
	Use:   "prune",
	Short: "Remove extraneous packages",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pruneCmd).Standalone()
	pruneCmd.Flags().Bool("dry-run", false, "only report changes")
	pruneCmd.Flags().Bool("json", false, "output as json")
	pruneCmd.Flags().StringArray("omit", []string{""}, "omit package type")
	addWorkspaceFlags(pruneCmd)

	rootCmd.AddCommand(pruneCmd)

	carapace.Gen(pruneCmd).PositionalAnyCompletion(
		action.ActionModules(),
	)
}
