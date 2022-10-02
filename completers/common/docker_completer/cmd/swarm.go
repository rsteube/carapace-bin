package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var swarmCmd = &cobra.Command{
	Use:   "swarm",
	Short: "Manage Swarm",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(swarmCmd).Standalone()

	rootCmd.AddCommand(swarmCmd)
}
