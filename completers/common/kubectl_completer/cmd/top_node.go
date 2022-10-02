package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/kubectl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var top_nodeCmd = &cobra.Command{
	Use:   "node",
	Short: "Display Resource (CPU/Memory/Storage) usage of nodes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(top_nodeCmd).Standalone()

	top_nodeCmd.Flags().String("heapster-namespace", "", "Namespace Heapster service is located in")
	top_nodeCmd.Flags().String("heapster-port", "", "Port name in service to use")
	top_nodeCmd.Flags().String("heapster-scheme", "", "Scheme (http or https) to connect to Heapster as")
	top_nodeCmd.Flags().String("heapster-service", "", "Name of Heapster service")
	top_nodeCmd.Flags().Bool("no-headers", false, "If present, print output without headers")
	top_nodeCmd.Flags().StringP("selector", "l", "", "Selector (label query) to filter on, supports '=', '==', and '!='.(e.g. -l key1=value1,key2=value2)")
	top_nodeCmd.Flags().String("sort-by", "", "If non-empty, sort nodes list using specified field. The field can be either 'cpu' or 'memory'.")
	topCmd.AddCommand(top_nodeCmd)

	carapace.Gen(top_nodeCmd).FlagCompletion(carapace.ActionMap{
		"heapster-namespace": action.ActionResources("", "namespaces"),
		"heapster-scheme":    carapace.ActionValues("http", "https"),
		"sort-by":            carapace.ActionValues("cpu", "memory"),
	})

	carapace.Gen(top_nodeCmd).PositionalCompletion(
		action.ActionResources("", "nodes"),
	)
}
