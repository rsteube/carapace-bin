package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var orb_publish_promoteCmd = &cobra.Command{
	Use:   "promote",
	Short: "Promote a development version of an orb to a semantic release",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(orb_publish_promoteCmd).Standalone()
	orb_publishCmd.AddCommand(orb_publish_promoteCmd)
}
