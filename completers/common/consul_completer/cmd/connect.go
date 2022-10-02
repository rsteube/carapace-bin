package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "Interact with Consul Connect",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connectCmd).Standalone()

	rootCmd.AddCommand(connectCmd)
}
