package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var create_tenantCmd = &cobra.Command{
	Use:   "tenant",
	Short: "Create or update a tenant",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(create_tenantCmd).Standalone()
	create_tenantCmd.Flags().String("cluster-role", "cluster-admin", "cluster role of the tenant role binding")
	create_tenantCmd.Flags().StringSlice("with-namespace", []string{}, "namespace belonging to this tenant")
	createCmd.AddCommand(create_tenantCmd)
}
