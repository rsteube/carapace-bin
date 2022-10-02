package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var domains_lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "Show all domains in a list",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(domains_lsCmd).Standalone()

	domainsCmd.AddCommand(domains_lsCmd)
}
