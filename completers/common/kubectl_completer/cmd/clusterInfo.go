package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var clusterInfoCmd = &cobra.Command{
	Use:   "cluster-info",
	Short: "Display cluster info",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(clusterInfoCmd).Standalone()

	rootCmd.AddCommand(clusterInfoCmd)
}
