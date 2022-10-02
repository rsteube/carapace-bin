package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var node_lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List nodes in the swarm",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(node_lsCmd).Standalone()

	node_lsCmd.Flags().StringP("filter", "f", "", "Filter output based on conditions provided")
	node_lsCmd.Flags().String("format", "", "Pretty-print nodes using a Go template")
	node_lsCmd.Flags().BoolP("quiet", "q", false, "Only display IDs")
	nodeCmd.AddCommand(node_lsCmd)
}
