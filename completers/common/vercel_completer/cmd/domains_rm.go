package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var domains_rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove a domain",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(domains_rmCmd).Standalone()

	domainsCmd.AddCommand(domains_rmCmd)

	// TODO positional completion
}
