package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var dns_lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List all DNS entries for a domain",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dns_lsCmd).Standalone()

	dnsCmd.AddCommand(dns_lsCmd)
}
