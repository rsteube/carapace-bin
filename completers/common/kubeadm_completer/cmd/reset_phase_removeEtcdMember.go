package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var reset_phase_removeEtcdMemberCmd = &cobra.Command{
	Use:   "remove-etcd-member",
	Short: "Remove a local etcd member.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(reset_phase_removeEtcdMemberCmd).Standalone()
	reset_phase_removeEtcdMemberCmd.Flags().String("kubeconfig", "/etc/kubernetes/admin.conf", "The kubeconfig file to use when talking to the cluster. If the flag is not set, a set of standard locations can be searched for an existing kubeconfig file.")
	reset_phaseCmd.AddCommand(reset_phase_removeEtcdMemberCmd)

	carapace.Gen(reset_phase_removeEtcdMemberCmd).FlagCompletion(carapace.ActionMap{
		"kubeconfig": carapace.ActionFiles(),
	})
}
