package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var init_phase_kubeletFinalizeCmd = &cobra.Command{
	Use:   "kubelet-finalize",
	Short: "Updates settings relevant to the kubelet after TLS bootstrap",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(init_phase_kubeletFinalizeCmd).Standalone()
	init_phaseCmd.AddCommand(init_phase_kubeletFinalizeCmd)
}
