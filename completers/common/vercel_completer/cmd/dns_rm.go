package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var dns_rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove a DNS entry using its ID",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dns_rmCmd).Standalone()

	dnsCmd.AddCommand(dns_rmCmd)

	// TODO positional completion
}
