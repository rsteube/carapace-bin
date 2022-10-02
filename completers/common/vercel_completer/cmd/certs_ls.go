package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var certs_lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "Show all available certificates",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(certs_lsCmd).Standalone()

	certsCmd.AddCommand(certs_lsCmd)
}
