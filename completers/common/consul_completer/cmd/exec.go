package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/consul_completer/cmd/action"
	"github.com/spf13/cobra"
)

var execCmd = &cobra.Command{
	Use:   "exec",
	Short: "Executes a command on Consul nodes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(execCmd).Standalone()
	addClientFlags(execCmd)
	addServerFlags(execCmd)

	execCmd.Flags().String("node", "", "Regular expression to filter on node names.")
	execCmd.Flags().String("prefix", "", "Prefix in the KV store to use for request data.")
	execCmd.Flags().String("service", "", "Regular expression to filter on service instances.")
	execCmd.Flags().Bool("shell", false, "Use a shell to run the command.")
	execCmd.Flags().String("tag", "", "Regular expression to filter on service tags.")
	execCmd.Flags().Bool("verbose", false, "Enables verbose output.")
	execCmd.Flags().String("wait", "", "Period to wait with no responses before terminating execution.")
	execCmd.Flags().String("wait-repl", "", "Period to wait for replication before firing event.")
	rootCmd.AddCommand(execCmd)

	carapace.Gen(execCmd).FlagCompletion(carapace.ActionMap{
		"node": action.ActionNodes(execCmd),
	})
}
