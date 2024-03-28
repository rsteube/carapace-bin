package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var address_showdumpCmd = &cobra.Command{
	Use:   "showdump",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(address_showdumpCmd).Standalone()

	addressCmd.AddCommand(address_showdumpCmd)
}
