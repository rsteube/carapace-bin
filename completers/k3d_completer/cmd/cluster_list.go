package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var cluster_listCmd = &cobra.Command{
	Use:     "list [NAME [NAME...]]",
	Short:   "List cluster(s)",
	Aliases: []string{"ls", "get"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cluster_listCmd).Standalone()

	cluster_listCmd.Flags().Bool("no-headers", false, "Disable headers")
	cluster_listCmd.Flags().StringP("output", "o", "", "Output format. One of: json|yaml")
	cluster_listCmd.Flags().Bool("token", false, "Print k3s cluster token")
	clusterCmd.AddCommand(cluster_listCmd)
}
