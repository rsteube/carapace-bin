package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var certs_rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove a certificate by id",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(certs_rmCmd).Standalone()

	certsCmd.AddCommand(certs_rmCmd)

	// TODO positional completion
}
