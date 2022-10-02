package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var authMethod_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all configured auth methods",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(authMethod_listCmd).Standalone()

	authMethodCmd.AddCommand(authMethod_listCmd)
}
