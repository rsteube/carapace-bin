package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var registry_syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Synchronizes all registries",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(registry_syncCmd).Standalone()
	registryCmd.AddCommand(registry_syncCmd)
}
