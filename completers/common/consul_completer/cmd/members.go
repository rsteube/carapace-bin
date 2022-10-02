package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var membersCmd = &cobra.Command{
	Use:   "members",
	Short: "Lists the members of a Consul cluster",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(membersCmd).Standalone()
	addClientFlags(membersCmd)

	membersCmd.Flags().Bool("detailed", false, "Provides detailed information about nodes.")
	membersCmd.Flags().String("segment", "", "(Enterprise-only) If provided, output is filtered to only nodes inthe given segment.")
	membersCmd.Flags().String("status", "", "If provided, output is filtered to only nodes matching the regular expression for status.")
	membersCmd.Flags().Bool("wan", false, "If the agent is in server mode, this can be used to return the other peers in the WAN pool.")
	rootCmd.AddCommand(membersCmd)

	carapace.Gen(membersCmd).FlagCompletion(carapace.ActionMap{
		"status": carapace.ActionValues("alive", "left", "failed"),
	})
}
