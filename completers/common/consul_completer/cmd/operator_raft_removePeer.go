package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/consul_completer/cmd/action"
	"github.com/spf13/cobra"
)

var operator_raft_removePeerCmd = &cobra.Command{
	Use:   "removePeer",
	Short: "Remove a Consul server from the Raft configuration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(operator_raft_removePeerCmd).Standalone()
	addClientFlags(operator_raft_removePeerCmd)
	addServerFlags(operator_raft_removePeerCmd)
	operator_raft_removePeerCmd.Flags().String("address", "", "The address to remove from the Raft configuration")
	operator_raft_removePeerCmd.Flags().String("id", "", "The ID to remove from the Raft configuration")

	operator_raftCmd.AddCommand(operator_raft_removePeerCmd)

	carapace.Gen(operator_raft_removePeerCmd).FlagCompletion(carapace.ActionMap{
		"address": action.ActionRaftPeerAddresses(operator_raft_removePeerCmd),
		"id":      action.ActionRaftPeerIds(operator_raft_removePeerCmd),
	})
}
