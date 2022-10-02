package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var dns_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new DNS entry",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dns_addCmd).Standalone()

	dnsCmd.AddCommand(dns_addCmd)

	// TODO positional completion
	carapace.Gen(dnsCmd).PositionalCompletion(
		carapace.ActionValues(),
		carapace.ActionValues(),
		carapace.ActionValues("A", "CAA", "AAAA", "ALIAS", "CNAME", "MS", "SRV", "TXT"),
	)
}
