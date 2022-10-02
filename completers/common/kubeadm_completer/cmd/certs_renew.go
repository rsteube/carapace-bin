package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var certs_renewCmd = &cobra.Command{
	Use:   "renew",
	Short: "Renew certificates for a Kubernetes cluster",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(certs_renewCmd).Standalone()
	certsCmd.AddCommand(certs_renewCmd)
}
