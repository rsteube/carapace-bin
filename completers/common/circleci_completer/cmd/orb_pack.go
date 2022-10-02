package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var orb_packCmd = &cobra.Command{
	Use:   "pack",
	Short: "Pack an Orb with local scripts.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(orb_packCmd).Standalone()
	orbCmd.AddCommand(orb_packCmd)
}
