package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var connection_loadCmd = &cobra.Command{
	Use:   "load",
	Short: "Load or reload one or more connection files from disk",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connection_loadCmd).Standalone()

	connectionCmd.AddCommand(connection_loadCmd)

	carapace.Gen(connection_loadCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
