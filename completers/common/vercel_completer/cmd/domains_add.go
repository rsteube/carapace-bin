package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var domains_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new domain that you already own",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(domains_addCmd).Standalone()

	domainsCmd.AddCommand(domains_addCmd)
}
