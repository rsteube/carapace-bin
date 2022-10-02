package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var secrets_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new secret",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(secrets_addCmd).Standalone()

	secretsCmd.AddCommand(secrets_addCmd)
}
