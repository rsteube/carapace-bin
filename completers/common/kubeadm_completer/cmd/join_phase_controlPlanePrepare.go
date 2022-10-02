package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var join_phase_controlPlanePrepareCmd = &cobra.Command{
	Use:   "control-plane-prepare",
	Short: "Prepare the machine for serving a control plane",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(join_phase_controlPlanePrepareCmd).Standalone()
	join_phaseCmd.AddCommand(join_phase_controlPlanePrepareCmd)
}
