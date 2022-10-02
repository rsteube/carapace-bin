package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var orb_publish_incrementCmd = &cobra.Command{
	Use:   "increment",
	Short: "Increment a released version of an orb",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(orb_publish_incrementCmd).Standalone()
	orb_publishCmd.AddCommand(orb_publish_incrementCmd)
}
