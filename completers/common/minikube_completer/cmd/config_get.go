package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/minikube_completer/cmd/action"
	"github.com/spf13/cobra"
)

var config_getCmd = &cobra.Command{
	Use:   "get",
	Short: "Gets the value of PROPERTY_NAME from the minikube config file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	configCmd.AddCommand(config_getCmd)

	carapace.Gen(config_getCmd).PositionalCompletion(
		action.ActionConfigNames(),
	)
}
