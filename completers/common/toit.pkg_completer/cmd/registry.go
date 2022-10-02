package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var registryCmd = &cobra.Command{
	Use:   "registry",
	Short: "Manages registries",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(registryCmd).Standalone()
	rootCmd.AddCommand(registryCmd)
}
