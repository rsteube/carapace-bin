package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var domains_buyCmd = &cobra.Command{
	Use:   "buy",
	Short: "Buy a domain that you don't yet own",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(domains_buyCmd).Standalone()

	domainsCmd.AddCommand(domains_buyCmd)

	// TODO positional completion
}
