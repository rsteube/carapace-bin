package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/consul_completer/cmd/action"
	"github.com/spf13/cobra"
)

var services_deregisterCmd = &cobra.Command{
	Use:   "deregister",
	Short: "Deregister services with the local agent",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(services_deregisterCmd).Standalone()
	addClientFlags(services_deregisterCmd)
	addServerFlags(services_deregisterCmd)

	services_deregisterCmd.Flags().String("id", "", "ID to delete.")
	services_deregisterCmd.Flags().String("namespace", "", "Specifies the namespace to query.")
	servicesCmd.AddCommand(services_deregisterCmd)

	// TODO flag completion
	carapace.Gen(services_deregisterCmd).FlagCompletion(carapace.ActionMap{
		"id": action.ActionServices(services_deregisterCmd),
	})

	carapace.Gen(services_deregisterCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
