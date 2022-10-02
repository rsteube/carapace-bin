package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var domains_transferInCmd = &cobra.Command{
	Use:   "transfer-in",
	Short: "Transfer in a domain to Vercel",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(domains_transferInCmd).Standalone()

	domainsCmd.AddCommand(domains_transferInCmd)

	// TODO positional completion
}
