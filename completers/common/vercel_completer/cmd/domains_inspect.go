package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var domains_inspectCmd = &cobra.Command{
	Use:   "inspect",
	Short: "Displays information related to a domain",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(domains_inspectCmd).Standalone()

	domainsCmd.AddCommand(domains_inspectCmd)

	// TODO positional completion
}
