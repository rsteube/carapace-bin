package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var trust_signerCmd = &cobra.Command{
	Use:   "signer",
	Short: "Manage entities who can sign Docker images",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(trust_signerCmd).Standalone()

	trustCmd.AddCommand(trust_signerCmd)
}
