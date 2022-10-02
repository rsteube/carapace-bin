package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var orb_publishCmd = &cobra.Command{
	Use:   "publish",
	Short: "Publish an orb to the registry",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(orb_publishCmd).Standalone()
	orbCmd.AddCommand(orb_publishCmd)
}
