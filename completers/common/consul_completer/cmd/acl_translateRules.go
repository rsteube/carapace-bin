package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var acl_translateRulesCmd = &cobra.Command{
	Use:   "translate-rules",
	Short: "Translate the legacy rule syntax into the current syntax",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(acl_translateRulesCmd).Standalone()
	addClientFlags(acl_translateRulesCmd)
	addServerFlags(acl_translateRulesCmd)

	acl_translateRulesCmd.Flags().Bool("token-accessor", false, "Specifies that the TRANSLATE argument refers to a ACL token AccessorID.")
	acl_translateRulesCmd.Flags().Bool("token-secret", false, "Specifies that the TRANSLATE argument refers to a ACL token SecretID.")
	aclCmd.AddCommand(acl_translateRulesCmd)
}
