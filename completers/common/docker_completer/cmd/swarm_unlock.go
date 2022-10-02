package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var swarm_unlockCmd = &cobra.Command{
	Use:   "unlock",
	Short: "Unlock swarm",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(swarm_unlockCmd).Standalone()

	swarmCmd.AddCommand(swarm_unlockCmd)
}
