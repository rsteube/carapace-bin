package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var certs_certificateKeyCmd = &cobra.Command{
	Use:   "certificate-key",
	Short: "Generate certificate keys",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(certs_certificateKeyCmd).Standalone()
	certsCmd.AddCommand(certs_certificateKeyCmd)
}
