package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var registry_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List registries",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(registry_listCmd).Standalone()
	registryCmd.AddCommand(registry_listCmd)
}
