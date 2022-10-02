package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/minikube_completer/cmd/action"
	"github.com/spf13/cobra"
)

var ipCmd = &cobra.Command{
	Use:   "ip",
	Short: "Retrieves the IP address of the specified node",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	ipCmd.Flags().StringP("node", "n", "", "The node to get IP. Defaults to the primary control plane.")
	rootCmd.AddCommand(ipCmd)

	carapace.Gen(ipCmd).FlagCompletion(carapace.ActionMap{
		"node": action.ActionNodes(),
	})
}
