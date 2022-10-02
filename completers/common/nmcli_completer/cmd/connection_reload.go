package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var connection_reloadCmd = &cobra.Command{
	Use:   "reload",
	Short: "Reload all connection files from disk",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connection_reloadCmd).Standalone()

	connectionCmd.AddCommand(connection_reloadCmd)
}
