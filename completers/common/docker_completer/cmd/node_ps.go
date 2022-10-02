package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var node_psCmd = &cobra.Command{
	Use:   "ps",
	Short: "List tasks running on one or more nodes, defaults to current node",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(node_psCmd).Standalone()

	node_psCmd.Flags().StringP("filter", "f", "", "Filter output based on conditions provided")
	node_psCmd.Flags().String("format", "", "Pretty-print tasks using a Go template")
	node_psCmd.Flags().Bool("no-resolve", false, "Do not map IDs to Names")
	node_psCmd.Flags().Bool("no-trunc", false, "Do not truncate output")
	node_psCmd.Flags().BoolP("quiet", "q", false, "Only display task IDs")
	nodeCmd.AddCommand(node_psCmd)

	carapace.Gen(node_psCmd).PositionalAnyCompletion(docker.ActionNodes())
}
