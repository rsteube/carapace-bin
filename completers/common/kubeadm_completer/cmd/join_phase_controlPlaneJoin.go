package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var join_phase_controlPlaneJoinCmd = &cobra.Command{
	Use:   "control-plane-join",
	Short: "Join a machine as a control plane instance",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(join_phase_controlPlaneJoinCmd).Standalone()
	join_phaseCmd.AddCommand(join_phase_controlPlaneJoinCmd)
}
