package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var dnsCmd = &cobra.Command{
	Use:   "dns",
	Short: "Manages your DNS records",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dnsCmd).Standalone()

	rootCmd.AddCommand(dnsCmd)
}
