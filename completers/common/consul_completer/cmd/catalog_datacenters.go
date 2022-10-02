package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var catalog_datacentersCmd = &cobra.Command{
	Use:   "datacenters",
	Short: "Lists all known datacenters",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(catalog_datacentersCmd).Standalone()
	addClientFlags(catalog_datacentersCmd)
	addServerFlags(catalog_datacentersCmd)

	catalogCmd.AddCommand(catalog_datacentersCmd)
}
