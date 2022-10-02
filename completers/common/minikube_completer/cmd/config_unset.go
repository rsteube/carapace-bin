package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/minikube_completer/cmd/action"
	"github.com/spf13/cobra"
)

var config_unsetCmd = &cobra.Command{
	Use:   "unset",
	Short: "unsets an individual value in a minikube config file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	configCmd.AddCommand(config_unsetCmd)

	carapace.Gen(config_unsetCmd).PositionalCompletion(
		action.ActionConfigNames(),
	)
}
