package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var tls_certCmd = &cobra.Command{
	Use:   "cert",
	Short: "Helpers for certificates",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tls_certCmd).Standalone()

	tlsCmd.AddCommand(tls_certCmd)
}
