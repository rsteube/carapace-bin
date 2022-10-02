package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/npm"
	"github.com/spf13/cobra"
)

var explainCmd = &cobra.Command{
	Use:   "explain",
	Short: "Explain installed packages",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(explainCmd).Standalone()
	explainCmd.Flags().Bool("json", false, "output as json")
	addWorkspaceFlags(explainCmd)

	rootCmd.AddCommand(explainCmd)

	carapace.Gen(explainCmd).PositionalCompletion(
		npm.ActionDependencies(),
	)
}
