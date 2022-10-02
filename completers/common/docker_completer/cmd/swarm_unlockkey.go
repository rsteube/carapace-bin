package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var swarm_unlockkeyCmd = &cobra.Command{
	Use:   "unlock-key",
	Short: "Manage the unlock key",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(swarm_unlockkeyCmd).Standalone()

	swarm_unlockkeyCmd.Flags().BoolP("quiet", "q", false, "Only display token")
	swarm_unlockkeyCmd.Flags().Bool("rotate", false, "Rotate unlock key")
	swarmCmd.AddCommand(swarm_unlockkeyCmd)
}
