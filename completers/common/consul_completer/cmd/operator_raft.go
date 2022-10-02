package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var operator_raftCmd = &cobra.Command{
	Use:   "raft",
	Short: "Provides cluster-level tools for Consul operators",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(operator_raftCmd).Standalone()

	operatorCmd.AddCommand(operator_raftCmd)
}
