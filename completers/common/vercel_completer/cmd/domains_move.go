package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var domains_moveCmd = &cobra.Command{
	Use:   "move",
	Short: "Move a domain to another user or team",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(domains_moveCmd).Standalone()

	domainsCmd.AddCommand(domains_moveCmd)

	// TODO positional completion
}
