package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var swarm_jointokenCmd = &cobra.Command{
	Use:   "join-token",
	Short: "Manage join tokens",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(swarm_jointokenCmd).Standalone()

	swarm_jointokenCmd.Flags().BoolP("quiet", "q", false, "Only display token")
	swarm_jointokenCmd.Flags().Bool("rotate", false, "Rotate join token")
	swarmCmd.AddCommand(swarm_jointokenCmd)

	carapace.Gen(swarm_jointokenCmd).PositionalCompletion(
		carapace.ActionValues("worker", "manager"),
	)
}
