package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var mdnsCmd = &cobra.Command{
	Use:   "mdns",
	Short: "manage mdns",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mdnsCmd).Standalone()

	rootCmd.AddCommand(mdnsCmd)
}
