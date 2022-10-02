package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var authMethod_set_oidcCmd = &cobra.Command{
	Use:   "oidc",
	Short: "Configure an OIDC auth method",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(authMethod_set_oidcCmd).Standalone()

	authMethod_set_oidcCmd.Flags().String("access-selector", "", "Selector expression to control access based on claims.")
	authMethod_set_oidcCmd.Flags().String("description", "", "Short description of this auth method")
	authMethod_set_oidcCmd.Flags().String("display-name", "", "Display name for the UI")

	addGlobalOptions(authMethod_set_oidcCmd)
	addOidcAuthMethodOptions(authMethod_set_oidcCmd)

	authMethod_setCmd.AddCommand(authMethod_set_oidcCmd)
}
