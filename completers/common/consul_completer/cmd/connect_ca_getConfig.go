package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var connect_ca_getConfigCmd = &cobra.Command{
	Use:   "get-config",
	Short: "Display the current Connect Certificate Authority (CA) configuration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_ca_getConfigCmd).Standalone()
	addClientFlags(connect_ca_getConfigCmd)
	addServerFlags(connect_ca_getConfigCmd)

	connect_caCmd.AddCommand(connect_ca_getConfigCmd)
}
