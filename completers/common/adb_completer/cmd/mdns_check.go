package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var mdns_checkCmd = &cobra.Command{
	Use:   "check",
	Short: "check if mdns discovery is available",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mdns_checkCmd).Standalone()

	mdnsCmd.AddCommand(mdns_checkCmd)
}
