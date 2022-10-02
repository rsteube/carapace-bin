package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/npm_completer/cmd/action"
	"github.com/spf13/cobra"
)

var viewCmd = &cobra.Command{
	Use:   "view",
	Short: "View registry info",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(viewCmd).Standalone()
	viewCmd.Flags().Bool("json", false, "output as json")
	addWorkspaceFlags(viewCmd)

	rootCmd.AddCommand(viewCmd)

	carapace.Gen(viewCmd).PositionalCompletion(
		action.ActionPackages(viewCmd),
		// TODO field completion
	)
}
