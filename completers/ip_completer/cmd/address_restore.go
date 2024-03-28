package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var address_restoreCmd = &cobra.Command{
	Use:   "restore",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(address_restoreCmd).Standalone()

	addressCmd.AddCommand(address_restoreCmd)
}
