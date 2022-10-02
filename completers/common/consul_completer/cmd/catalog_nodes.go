package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/consul_completer/cmd/action"
	"github.com/spf13/cobra"
)

var catalog_nodesCmd = &cobra.Command{
	Use:   "nodes",
	Short: "Lists all nodes in the given datacenter",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(catalog_nodesCmd).Standalone()
	addClientFlags(catalog_nodesCmd)
	addServerFlags(catalog_nodesCmd)

	catalog_nodesCmd.Flags().Bool("detailed", false, "Output detailed information about the nodes including their addresses and metadata.")
	catalog_nodesCmd.Flags().String("filter", "", "Filter to use with the request")
	catalog_nodesCmd.Flags().String("near", "", "Node name to sort the node list in ascending order based on estimated round-trip time from that node.")
	catalog_nodesCmd.Flags().String("node-meta", "", "Metadata to filter nodes with the given `key=value` pairs.")
	catalog_nodesCmd.Flags().String("service", "", "Service `id or name` to filter nodes.")
	catalogCmd.AddCommand(catalog_nodesCmd)

	carapace.Gen(catalog_nodesCmd).FlagCompletion(carapace.ActionMap{
		"near":    action.ActionNodes(catalog_nodesCmd),
		"service": action.ActionServices(catalog_nodesCmd),
	})
}
