package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var network_pruneCmd = &cobra.Command{
	Use:   "prune",
	Short: "Remove all unused networks",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(network_pruneCmd).Standalone()

	network_pruneCmd.Flags().String("filter", "", "Provide filter values (e.g. 'until=<timestamp>')")
	network_pruneCmd.Flags().BoolP("force", "f", false, "Do not prompt for confirmation")
	networkCmd.AddCommand(network_pruneCmd)
}
