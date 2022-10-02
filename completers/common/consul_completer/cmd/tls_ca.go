package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var tls_caCmd = &cobra.Command{
	Use:   "ca",
	Short: "Helpers for CAs",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tls_caCmd).Standalone()

	tlsCmd.AddCommand(tls_caCmd)
}
