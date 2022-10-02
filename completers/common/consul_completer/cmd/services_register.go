package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var services_registerCmd = &cobra.Command{
	Use:   "register",
	Short: "Register services with the local agent",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(services_registerCmd).Standalone()
	addClientFlags(services_registerCmd)
	addServerFlags(services_registerCmd)

	services_registerCmd.Flags().String("address", "", "Address of the service to register for arg-based registration.")
	services_registerCmd.Flags().String("id", "", "ID of the service to register for arg-based registration.")
	services_registerCmd.Flags().String("kind", "", "The services 'kind'")
	services_registerCmd.Flags().String("meta", "", "Metadata to set on the service, formatted as key=value.")
	services_registerCmd.Flags().String("name", "", "Name of the service to register for arg-based registration.")
	services_registerCmd.Flags().String("namespace", "", "Specifies the namespace to query.")
	services_registerCmd.Flags().String("port", "", "Port of the service to register for arg-based registration.")
	services_registerCmd.Flags().String("tag", "", "Tag to add to the service.")
	services_registerCmd.Flags().String("tagged-address", "", "Tagged address to set on the service, formatted as key=value.")
	servicesCmd.AddCommand(services_registerCmd)

	// TODO flag completion

	carapace.Gen(services_registerCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
