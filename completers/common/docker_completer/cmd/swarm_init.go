package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var swarm_initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a swarm",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(swarm_initCmd).Standalone()

	swarm_initCmd.Flags().String("advertise-addr", "", "Advertised address (format: <ip|interface>[:port])")
	swarm_initCmd.Flags().Bool("autolock", false, "Enable manager autolocking (requiring an unlock")
	swarm_initCmd.Flags().String("availability", "", "Availability of the node")
	swarm_initCmd.Flags().String("cert-expiry", "", "Validity period for node certificates")
	swarm_initCmd.Flags().String("data-path-addr", "", "Address or interface to use for data path")
	swarm_initCmd.Flags().String("data-path-port", "", "Port number to use for data path traffic (1024 -")
	swarm_initCmd.Flags().String("default-addr-pool", "", "default address pool in CIDR format (default [])")
	swarm_initCmd.Flags().String("default-addr-pool-mask-length", "", "default address pool subnet mask length (default 24)")
	swarm_initCmd.Flags().String("dispatcher-heartbeat", "", "Dispatcher heartbeat period (ns|us|ms|s|m|h)")
	swarm_initCmd.Flags().String("external-ca", "", "Specifications of one or more certificate")
	swarm_initCmd.Flags().Bool("force-new-cluster", false, "Force create a new cluster from current state")
	swarm_initCmd.Flags().String("listen-addr", "", "Listen address (format: <ip|interface>[:port])")
	swarm_initCmd.Flags().String("max-snapshots", "", "Number of additional Raft snapshots to retain")
	swarm_initCmd.Flags().String("snapshot-interval", "", "Number of log entries between Raft snapshots")
	swarm_initCmd.Flags().String("task-history-limit", "", "Task history retention limit (default 5)")
	swarmCmd.AddCommand(swarm_initCmd)

	carapace.Gen(swarm_initCmd).FlagCompletion(carapace.ActionMap{
		"availability": carapace.ActionValues("active", "pause", "drain"),
	})
}
