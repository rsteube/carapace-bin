package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/direnv"
	"github.com/spf13/cobra"
)

var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Opens PATH_TO_RC or the current .envrc or .env into an $EDITOR and allow the file to be loaded afterwards",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(editCmd).Standalone()

	rootCmd.AddCommand(editCmd)

	carapace.Gen(editCmd).PositionalCompletion(
		carapace.Batch(
			direnv.ActionAuths(),
			carapace.ActionFiles(),
		).ToA(),
	)
}
