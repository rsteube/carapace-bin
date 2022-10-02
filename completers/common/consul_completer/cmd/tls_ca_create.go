package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var tls_ca_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new consul CA",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tls_ca_createCmd).Standalone()

	tls_ca_createCmd.Flags().String("additional-name-constraint", "", "Add name constraints for the CA.")
	tls_ca_createCmd.Flags().String("cluster-id", "", "ClusterID of the consul cluster.")
	tls_ca_createCmd.Flags().String("common-name", "", "Common Name of CA.")
	tls_ca_createCmd.Flags().String("days", "", "Provide number of days the CA is valid for from now on.")
	tls_ca_createCmd.Flags().String("domain", "", "Domain of consul cluster.")
	tls_ca_createCmd.Flags().Bool("name-constraint", false, "Add name constraints for the CA.")
	tls_caCmd.AddCommand(tls_ca_createCmd)

	// TODO flag completion
}
