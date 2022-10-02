package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/consul_completer/cmd/action"
	"github.com/spf13/cobra"
)

var catalog_servicesCmd = &cobra.Command{
	Use:   "services",
	Short: "Lists all registered services in a datacenter",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(catalog_servicesCmd).Standalone()
	addClientFlags(catalog_servicesCmd)
	addServerFlags(catalog_servicesCmd)

	catalog_servicesCmd.Flags().String("namespace", "", "Specifies the namespace to query.")
	catalog_servicesCmd.Flags().String("node", "", "Node `id or name` for which to list services.")
	catalog_servicesCmd.Flags().String("node-meta", "", "Metadata to filter nodes with the given `key=value` pairs.")
	catalog_servicesCmd.Flags().Bool("tags", false, "Display each service's tags as a comma-separated list beside each service entry.")
	catalogCmd.AddCommand(catalog_servicesCmd)

	carapace.Gen(catalog_servicesCmd).FlagCompletion(carapace.ActionMap{
		"node": action.ActionNodes(catalog_servicesCmd),
	})
}
