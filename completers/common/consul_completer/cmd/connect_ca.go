package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var connect_caCmd = &cobra.Command{
	Use:   "ca",
	Short: "Interact with the Consul Connect Certificate Authority (CA)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_caCmd).Standalone()

	connectCmd.AddCommand(connect_caCmd)
}
