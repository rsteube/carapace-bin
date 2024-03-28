package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var create_secretCmd = &cobra.Command{
	Use:   "secret",
	Short: "Create or update Kubernetes secrets",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(create_secretCmd).Standalone()
	createCmd.AddCommand(create_secretCmd)
}
