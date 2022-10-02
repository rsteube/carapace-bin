package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/kubectl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var auth_canICmd = &cobra.Command{
	Use:   "can-i",
	Short: "Check whether an action is allowed",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(auth_canICmd).Standalone()

	auth_canICmd.Flags().BoolP("all-namespaces", "A", false, "If true, check the specified action in all namespaces.")
	auth_canICmd.Flags().Bool("list", false, "If true, prints all allowed actions.")
	auth_canICmd.Flags().Bool("no-headers", false, "If true, prints allowed actions without headers")
	auth_canICmd.Flags().BoolP("quiet", "q", false, "If true, suppress output and just return the exit code.")
	auth_canICmd.Flags().String("subresource", "", "SubResource such as pod/log or deployment/scale")
	authCmd.AddCommand(auth_canICmd)

	carapace.Gen(auth_canICmd).PositionalCompletion(
		action.ActionResourceVerbs(),
		action.ActionApiResourceResources(),
	)
}
