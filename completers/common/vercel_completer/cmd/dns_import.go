package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var dns_importCmd = &cobra.Command{
	Use:   "import",
	Short: "Import a DNS zone file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dns_importCmd).Standalone()

	dnsCmd.AddCommand(dns_importCmd)

	// TODO positional completion
	carapace.Gen(dnsCmd).PositionalCompletion(
		carapace.ActionValues(),
		carapace.ActionFiles(),
	)
}
