package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var operator_autopilot_setConfigCmd = &cobra.Command{
	Use:   "set-config",
	Short: "Modify the current Autopilot configuration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(operator_autopilot_setConfigCmd).Standalone()
	addClientFlags(operator_autopilot_setConfigCmd)
	addServerFlags(operator_autopilot_setConfigCmd)

	operator_autopilot_setConfigCmd.Flags().String("cleanup-dead-servers", "", "Controls whether Consul will automatically remove dead servers when new ones are successfully added.")
	operator_autopilot_setConfigCmd.Flags().String("disable-upgrade-migration", "", "(Enterprise-only) Controls whether Consul will avoid promoting new servers until it can perform a migration.")
	operator_autopilot_setConfigCmd.Flags().String("last-contact-threshold", "", "Controls the maximum amount of time a server can go without contact from the leader before being considered unhealthy.")
	operator_autopilot_setConfigCmd.Flags().String("max-trailing-logs", "", "Controls the maximum number of log entries that a server can trail the leader by before being considered unhealthy.")
	operator_autopilot_setConfigCmd.Flags().String("min-quorum", "", "Sets the minimum number of servers required in a cluster before autopilot is allowed to prune dead servers.")
	operator_autopilot_setConfigCmd.Flags().String("redundancy-zone-tag", "", "(Enterprise-only) Controls the node_meta tag name used for separating servers into different redundancy zones.")
	operator_autopilot_setConfigCmd.Flags().String("server-stabilization-time", "", "Controls the minimum amount of time a server must be stable in the 'healthy' state before being added to the cluster.")
	operator_autopilot_setConfigCmd.Flags().String("upgrade-version-tag", "", "(Enterprise-only) The node_meta tag to use for version info when performing upgrade migrations.")
	operator_autopilotCmd.AddCommand(operator_autopilot_setConfigCmd)

	carapace.Gen(operator_autopilot_setConfigCmd).FlagCompletion(carapace.ActionMap{
		"cleanup-dead-servers":      carapace.ActionValues("true", "false"),
		"disable-upgrade-migration": carapace.ActionValues("true", "false"),
	})
}
