package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var runner_agentCmd = &cobra.Command{
	Use:   "agent",
	Short: "Run a runner for executing remote operations",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(runner_agentCmd).Standalone()

	runner_agentCmd.Flags().Bool("enable-dynamic-config", false, "Allow dynamic config to be created when an exec plugin is used.")
	runner_agentCmd.Flags().String("id", "", "If this is set, the runner will use the specified id.")
	runner_agentCmd.Flags().String("liveness-tcp-addr", "", "Open a TCP listener on this address when it is running.")
	runner_agentCmd.Flags().Bool("odr", false, "Indicates to the runner it's operating as an on-demand runner.")

	addGlobalOptions(runner_agentCmd)

	runnerCmd.AddCommand(runner_agentCmd)
}
