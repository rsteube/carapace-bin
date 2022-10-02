package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var orb_unlistCmd = &cobra.Command{
	Use:   "unlist",
	Short: "Disable or enable an orb's listing in the registry",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(orb_unlistCmd).Standalone()
	orbCmd.AddCommand(orb_unlistCmd)
}
