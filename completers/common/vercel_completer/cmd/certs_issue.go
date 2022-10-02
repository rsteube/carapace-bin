package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var certs_issueCmd = &cobra.Command{
	Use:   "issue",
	Short: "Issue a new certificate for a domain",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(certs_issueCmd).Standalone()

	certsCmd.AddCommand(certs_issueCmd)

	// TODO positional completion
}
