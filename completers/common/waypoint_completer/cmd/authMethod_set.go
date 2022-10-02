package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var authMethod_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Create or update an auth method",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(authMethod_setCmd).Standalone()

	authMethodCmd.AddCommand(authMethod_setCmd)
}
