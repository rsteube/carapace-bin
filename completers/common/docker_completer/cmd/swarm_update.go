package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var swarm_updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update the swarm",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(swarm_updateCmd).Standalone()

	swarm_updateCmd.Flags().Bool("autolock", false, "Change manager autolocking setting (true|false)")
	swarm_updateCmd.Flags().String("cert-expiry", "", "Validity period for node certificates (ns|us|ms|s|m|h) (default 2160h0m0s)")
	swarm_updateCmd.Flags().String("dispatcher-heartbeat", "", "Dispatcher heartbeat period (ns|us|ms|s|m|h) (default 5s)")
	swarm_updateCmd.Flags().String("external-ca", "", "Specifications of one or more certificate signing endpoints")
	swarm_updateCmd.Flags().String("max-snapshots", "", "Number of additional Raft snapshots to retain")
	swarm_updateCmd.Flags().String("snapshot-interval", "", "Number of log entries between Raft snapshots (default 10000)")
	swarm_updateCmd.Flags().String("task-history-limit", "", "Task history retention limit (default 5)")
	swarmCmd.AddCommand(swarm_updateCmd)
}
