package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var init_phase_controlPlaneCmd = &cobra.Command{
	Use:   "control-plane",
	Short: "Generate all static Pod manifest files necessary to establish the control plane",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(init_phase_controlPlaneCmd).Standalone()
	init_phaseCmd.AddCommand(init_phase_controlPlaneCmd)
}
