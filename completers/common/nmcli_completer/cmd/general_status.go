package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var general_statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show overall status of NetworkManager",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(general_statusCmd).Standalone()

	generalCmd.AddCommand(general_statusCmd)
}
