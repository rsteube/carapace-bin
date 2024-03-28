package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var registry_deleteCmd = &cobra.Command{
	Use:     "delete (NAME | --all)",
	Short:   "Delete registry/registries.",
	Aliases: []string{"del", "rm"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(registry_deleteCmd).Standalone()

	registry_deleteCmd.Flags().BoolP("all", "a", false, "Delete all existing registries")
	registryCmd.AddCommand(registry_deleteCmd)
}
