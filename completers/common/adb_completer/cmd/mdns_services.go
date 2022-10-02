package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var mdns_servicesCmd = &cobra.Command{
	Use:   "services",
	Short: "list all discovered services",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mdns_servicesCmd).Standalone()

	mdnsCmd.AddCommand(mdns_servicesCmd)
}
