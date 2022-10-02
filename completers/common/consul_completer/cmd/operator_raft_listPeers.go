package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var operator_raft_listPeersCmd = &cobra.Command{
	Use:   "list-peers",
	Short: "Display the current Raft peer configuration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(operator_raft_listPeersCmd).Standalone()
	addClientFlags(operator_raft_listPeersCmd)
	addServerFlags(operator_raft_listPeersCmd)

	operator_raftCmd.AddCommand(operator_raft_listPeersCmd)
}
