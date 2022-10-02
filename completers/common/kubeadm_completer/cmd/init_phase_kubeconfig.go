package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var init_phase_kubeconfigCmd = &cobra.Command{
	Use:   "kubeconfig",
	Short: "Generate all kubeconfig files necessary to establish the control plane and the admin kubeconfig file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(init_phase_kubeconfigCmd).Standalone()
	init_phaseCmd.AddCommand(init_phase_kubeconfigCmd)
}
