package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var reset_phase_updateClusterStatusCmd = &cobra.Command{
	Use:   "update-cluster-status",
	Short: "Remove this node from the ClusterStatus object (DEPRECATED).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(reset_phase_updateClusterStatusCmd).Standalone()
	reset_phaseCmd.AddCommand(reset_phase_updateClusterStatusCmd)
}
