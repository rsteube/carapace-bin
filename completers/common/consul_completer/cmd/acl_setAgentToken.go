package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var acl_setAgentTokenCmd = &cobra.Command{
	Use:   "set-agent-token",
	Short: "Assign tokens for the Consul Agent's usage",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(acl_setAgentTokenCmd).Standalone()
	addClientFlags(acl_setAgentTokenCmd)
	addServerFlags(acl_setAgentTokenCmd)

	aclCmd.AddCommand(acl_setAgentTokenCmd)

	carapace.Gen(acl_setAgentTokenCmd).PositionalCompletion(
		carapace.ActionValuesDescribed(
			"default", "The token for both internal agent operations and operations initiated by the HTTP and DNS interfaces",
			"agent", "The token that the agent will use for internal agent operations.",
			"master", "This sets the token that can be used to access the Agent APIs in the event that the ACL datacenter cannot be reached.",
			"replication", "This is the token that the agent will use for replication operations.",
		),
	)
}
