package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var clusterCmd = &cobra.Command{
	Use:   "cluster",
	Short: "Manage cluster(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(clusterCmd).Standalone()

	rootCmd.AddCommand(clusterCmd)
}
